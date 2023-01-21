package handler

import (
	"github.com/solost23/protopb/gen/go/protos/oss"
	"oss_service/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::user service
	oss.RegisterOssServer(config.Server, service.NewOSSService(config.Sl, config.MysqlConnect, config.RedisClient, config.KafkaProducer, config.MinioClient))
	return
}
