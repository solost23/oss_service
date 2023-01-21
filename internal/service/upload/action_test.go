package upload

import (
	"context"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/common"
	"github.com/solost23/protopb/gen/go/protos/oss"
	"oss_service/configs"
	"oss_service/internal/minio_storage"
	"oss_service/internal/models"
	"testing"
)

func TestAction_Deal(t *testing.T) {
	minioClient, err := minio_storage.NewMinio(&configs.MinioConf{
		EndPoint:        "localhost:9000",
		AccessKeyId:     "minioadmin",
		SecretAccessKey: "minioadmin",
		UserSsl:         false,
	})
	if err != nil {
		t.Error(err.Error())
	}
	mdb, _ := models.InitMysql(&configs.MySQLConf{
		DataSourceName:  "root:123@tcp(localhost:3306)/oss_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		MaxConnLifeTime: 100,
	})
	type arg struct {
		ctx     context.Context
		request *oss.UploadRequest
	}
	type want struct {
		err error
	}
	type test struct {
		arg  arg
		want want
	}
	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &oss.UploadRequest{
					Header: &common.RequestHeader{
						OperatorUid: 100,
						TraceId:     101,
					},
					Folder:     "image",
					Key:        "xxx",
					Data:       []byte("content"),
					UploadType: "image",
				},
			},
			want: want{
				err: nil,
			},
		},
	}
	action := NewActionWithCtx(context.Background())
	action.SetMinio(minioClient)
	action.SetMysql(mdb)
	for _, test := range tests {
		reply, err := action.Deal(test.arg.ctx, test.arg.request)
		if err != test.want.err {
			t.Error(err.Error())
		}
		fmt.Println(reply.GetUrl(), reply.GetKey())
	}
}
