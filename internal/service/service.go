package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/gookit/slog"
	"github.com/solost23/go_interface/gen_go/oss"
	"gorm.io/gorm"
	"oss_service/internal/service/upload"
)

type OSSService struct {
	sl            *slog.SugaredLogger
	mdb           *gorm.DB
	rdb           *redis.Client
	kafkaProducer sarama.SyncProducer
	oss.UnimplementedOssServer
}

func NewOSSService(sl *slog.SugaredLogger, mdb *gorm.DB, rdb *redis.Client, kafkaProducer sarama.SyncProducer) *OSSService {
	return &OSSService{
		sl:            sl,
		mdb:           mdb,
		rdb:           rdb,
		kafkaProducer: kafkaProducer,
	}
}

// upload
func (h *OSSService) Upload(ctx context.Context, request *oss.UploadRequest) (reply *oss.UploadResponse, err error) {
	action := upload.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetSl(h.sl)
	action.SetMysql(h.mdb)
	return action.Deal(ctx, request)
}
