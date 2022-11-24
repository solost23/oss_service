package handler

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/gookit/slog"
	"github.com/minio/minio-go"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Server        *grpc.Server
	Sl            *slog.SugaredLogger
	MysqlConnect  *gorm.DB
	RedisClient   *redis.Client
	KafkaProducer sarama.SyncProducer
	MinioClient   *minio.Client
}
