package minio_storage

import (
	"context"
	"github.com/minio/minio-go"
	"io"
	"net/url"
	"time"
)

func CreateBucket(_ context.Context, minio *minio.Client, bucketName string) error {
	exists, err := minio.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// 不存在，则创建
	err = minio.MakeBucket(bucketName, "")
	if err != nil {
		return err
	}
	return nil
}

// 上传文件
func FileUpload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, filePath string, contentType string) (err error) {
	_, err = minioClient.FPutObjectWithContext(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return err
}

// 上传流数据
func StreamUpload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, reader io.Reader, objectSize int64, contentType string) (err error) {
	_, err = minioClient.PutObjectWithContext(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	return err
}

// 下载文件
func FileDownload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, filePath string) (err error) {
	err = minioClient.FGetObjectWithContext(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	return err
}

// 生成对象的url
func GetFileUrl(_ context.Context, minioClient *minio.Client, bucketName string, objectName string, expire time.Duration, request url.Values) (url string, err error) {
	urlObject, err := minioClient.PresignedGetObject(bucketName, objectName, expire, request)
	if err != nil {
		return "", err
	}
	return urlObject.String(), nil
}
