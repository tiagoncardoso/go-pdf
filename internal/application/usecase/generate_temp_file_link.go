package usecase

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tiagoncardoso/go-pdf/config"
)

type GenerateTempFileLink struct {
	env *config.EnvConfig
}

func NewGenerateTempFileLink(env *config.EnvConfig) *GenerateTempFileLink {
	return &GenerateTempFileLink{
		env,
	}
}

func (g *GenerateTempFileLink) Execute(objectKey string) (string, error) {
	storageSession, err := session.NewSession(&aws.Config{
		Region:           aws.String(g.env.StorageRegion),
		Endpoint:         aws.String(g.env.StorageEndpoint),
		Credentials:      credentials.NewStaticCredentials(g.env.StorageAccessKey, g.env.StorageSecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}

	storageService := s3.New(storageSession)

	err = checkFileExists(storageService, g.env.StorageSpaceName, objectKey)
	if err != nil {
		return "", err
	}

	req, _ := storageService.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(g.env.StorageSpaceName),
		Key:    aws.String(objectKey),
	})

	urlStr, err := req.Presign(time.Duration(g.env.PdfLinkExpirationSeconds) * time.Second)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func checkFileExists(storageService *s3.S3, storageSpaceName string, objectKey string) error {
	_, err := storageService.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(storageSpaceName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			return errors.New("file not found in storage")
		}
	}

	return nil
}
