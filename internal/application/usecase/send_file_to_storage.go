package usecase

import (
	"fmt"

	"github.com/tiagoncardoso/go-pdf/config"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type SendFileToStorage struct {
	endpoint  string
	spaceName string
	accessKey string
	secretKey string
	region    string
}

func NewSendFileToStorage(env *config.EnvConfig) *SendFileToStorage {
	return &SendFileToStorage{
		endpoint:  env.StorageEndpoint,
		spaceName: env.StorageSpaceName,
		accessKey: env.StorageAccessKey,
		secretKey: env.StorageSecretKey,
		region:    env.StorageRegion,
	}
}

func (s *SendFileToStorage) Execute(filePath string) error {
	filePath = "./internal/output/rep_1755560063.pdf"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	storageSession, err := session.NewSession(&aws.Config{
		Region:           aws.String(s.region),
		Endpoint:         aws.String(s.endpoint),
		Credentials:      credentials.NewStaticCredentials(s.accessKey, s.secretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return err
	}

	storageService := s3.New(storageSession)

	objectKey := "analytics/" + file.Name()

	_, err = storageService.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.spaceName),
		Key:    aws.String(objectKey),
		Body:   file,
		ACL:    aws.String("public-read"), // or "private"
	})
	if err != nil {
		return err
	}

	fmt.Println("File uploaded successfully!")

	return nil
}
