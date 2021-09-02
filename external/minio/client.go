package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rysmaadit/go-template/common/errors"
	log "github.com/sirupsen/logrus"
)

type Client interface {
	Ping() error
}

type client struct {
	minioClient *minio.Client
	bucketName  string
}

func (c *client) Ping() error {
	ctx := context.Background()
	isExist, err := c.minioClient.BucketExists(ctx, c.bucketName)
	if err != nil {
		log.Warning("Error in checking the bucket")
		return err
	}

	if !isExist {
		return errors.New(fmt.Sprintf("bucket %s does not exist", c.bucketName))
	}
	return nil
}

func NewMinioClient(config ClientConfig) *client {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV2(config.AccessKey, config.SecretKey, ""),
		Secure: true,
		Region: config.Region,
	})
	if err != nil {
		log.Fatalf("unable to initiate minio client. %v", err)
	}
	return &client{
		minioClient: minioClient,
		bucketName:  config.BucketName,
	}
}
