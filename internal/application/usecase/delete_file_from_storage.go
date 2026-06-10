package usecase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tiagoncardoso/go-pdf/config"
)

type DeleteFileFromStorage struct {
	env *config.EnvConfig
}

func NewDeleteFileFromStorage(env *config.EnvConfig) *DeleteFileFromStorage {
	return &DeleteFileFromStorage{
		env,
	}
}

func (d *DeleteFileFromStorage) Execute(objectKey string) error {
	storageSession, err := session.NewSession(&aws.Config{
		Region:           aws.String(d.env.StorageRegion),
		Endpoint:         aws.String(d.env.StorageEndpoint),
		Credentials:      credentials.NewStaticCredentials(d.env.StorageAccessKey, d.env.StorageSecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return err
	}

	storageService := s3.New(storageSession)

	_, err = storageService.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(d.env.StorageSpaceName),
		Key:    aws.String(objectKey),
	})

	return err
}
