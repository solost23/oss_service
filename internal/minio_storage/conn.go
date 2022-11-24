package minio_storage

import (
	"github.com/minio/minio-go"
	"oss_service/configs"
)

func NewMinio(config *configs.MinioConf) (client *minio.Client, err error) {
	client, err = minio.New(config.EndPoint, config.SecretAccessKey, config.AccessKeyId, config.UserSsl)
	if err != nil {
		return nil, err
	}
	return client, nil
}
