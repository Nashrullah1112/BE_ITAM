package minio

import (
	"context"
	"time"

	"github.com/banggibima/be-itam/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Client(config *config.Config) (*minio.Client, error) {
	endpoint := config.Minio.Endpoint
	accessKeyID := config.Minio.AccessKeyID
	secretAccessKey := config.Minio.SecretAccessKey
	useSSL := config.Minio.UseSSL
	token := ""

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, token),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateBucket(client *minio.Client, name, region string) error {
	err := client.MakeBucket(context.Background(), name, minio.MakeBucketOptions{
		Region:        region,
		ObjectLocking: false,
	})
	if err != nil {
		exists, err := client.BucketExists(context.Background(), name)
		if err == nil && exists {
			return nil
		}
		return err
	}
	return nil
}

func RemoveBucket(client *minio.Client, name string) error {
	err := client.RemoveBucket(context.Background(), name)
	if err != nil {
		return err
	}
	return nil
}

func ObjectUpload(client *minio.Client, bucket, object, path, filetype string) error {
	file, err := client.FPutObject(context.Background(), bucket, object, path, minio.PutObjectOptions{
		ContentType: filetype,
	})
	if err != nil {
		return err
	}

	_ = file

	return nil
}

func ObjectDownload(client *minio.Client, bucket, object, path string) error {
	err := client.FGetObject(context.Background(), bucket, object, path, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ObjectRemove(client *minio.Client, bucket, object string) error {
	err := client.RemoveObject(context.Background(), bucket, object, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ObjectPresignedURL(client *minio.Client, bucket, object string) (interface{}, error) {
	expiry := time.Duration(604800) * time.Second

	presignedURL, err := client.PresignedGetObject(context.Background(), bucket, object, expiry, nil)
	if err != nil {
		return nil, err
	}

	return presignedURL.String(), nil
}
