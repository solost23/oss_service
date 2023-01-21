package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/gookit/slog"
	"github.com/minio/minio-go"
	"github.com/solost23/protopb/gen/go/protos/oss"
	"gorm.io/gorm"
	"oss_service/internal/service/upload"
)

type OSSService struct {
	sl            *slog.SugaredLogger
	mdb           *gorm.DB
	rdb           *redis.Client
	kafkaProducer sarama.SyncProducer
	minioClient   *minio.Client
	oss.UnimplementedOssServer
}

func NewOSSService(sl *slog.SugaredLogger, mdb *gorm.DB, rdb *redis.Client, kafkaProducer sarama.SyncProducer, minioClient *minio.Client) *OSSService {
	return &OSSService{
		sl:            sl,
		mdb:           mdb,
		rdb:           rdb,
		kafkaProducer: kafkaProducer,
		minioClient:   minioClient,
	}
}

// upload
func (h *OSSService) Upload(ctx context.Context, request *oss.UploadRequest) (reply *oss.UploadResponse, err error) {
	action := upload.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetSl(h.sl)
	action.SetMysql(h.mdb)
	action.SetMinio(h.minioClient)
	return action.Deal(ctx, request)
}
