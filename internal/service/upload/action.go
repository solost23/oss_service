package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/oss"
	"net/url"
	"oss_service/internal/minio_storage"
	"oss_service/internal/models"
	"oss_service/internal/service/base"
	"time"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *oss.UploadRequest) (reply *oss.UploadResponse, err error) {
	err = minio_storage.CreateBucket(ctx, a.GetMinioClient(), request.GetFolder())
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(request.GetData())
	err = minio_storage.StreamUpload(ctx, a.GetMinioClient(), request.GetFolder(), request.GetKey(), reader, reader.Size(), fmt.Sprintf("Application/%s", request.GetUploadType()))
	if err != nil {
		return nil, err
	}
	requestPrams := make(url.Values)
	fileUrl, err := minio_storage.GetFileUrl(ctx, a.GetMinioClient(), request.GetFolder(), request.GetKey(), 168*time.Hour, requestPrams)
	if err != nil {
		return nil, err
	}
	reply = &oss.UploadResponse{
		Url: fileUrl,
		Key: request.GetKey(),
	}
	// 记录日志
	err = (&models.OSSFile{
		CreatorBase: models.CreatorBase{
			CreatorId: uint(request.GetHeader().GetOperatorUid()),
		},
		SaveName: request.GetKey(),
		SavePath: request.GetFolder(),
		FileType: request.GetUploadType(),
	}).Insert(a.GetMysqlConnect())
	return reply, nil
}
