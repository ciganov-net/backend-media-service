package storage

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appconfig "github.com/ciganov-net/backend-media-service/internal/config"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
}

func NewS3Storage(cfg appconfig.Config) *S3Storage {
	sdkConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.S3AccessKey,
				cfg.S3SecretKey,
				"",
			),
		),
	)

	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.S3Endpoint)
		o.UsePathStyle = *aws.Bool(true)
	})

	storage := &S3Storage{
		Client: client,
		Bucket: cfg.S3Bucket,
	}

	storage.ensureBucket()

	return storage
}

func (s *S3Storage) ensureBucket() {
	_, err := s.Client.HeadBucket(
		context.Background(),
		&s3.HeadBucketInput{
			Bucket: aws.String(s.Bucket),
		},
	)

	if err == nil {
		log.Println("bucket already exists")
		return
	}

	_, err = s.Client.CreateBucket(
		context.Background(),
		&s3.CreateBucketInput{
			Bucket: aws.String(s.Bucket),
		},
	)

	if err != nil {
		panic(err)
	}

	log.Println("bucket created")
}

func (s *S3Storage) Upload(
	objectKey string,
	body io.Reader,
	contentType string,
) error {
	_, err := s.Client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket:      aws.String(s.Bucket),
			Key:         aws.String(objectKey),
			Body:        body,
			ContentType: aws.String(contentType),
		},
	)

	return err
}

func (s *S3Storage) Delete(objectKey string) error {
	_, err := s.Client.DeleteObject(
		context.Background(),
		&s3.DeleteObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(objectKey),
		},
	)

	return err
}

func (s *S3Storage) GetPresignedURL(objectKey string) (string, error) {
	presignedClient := s3.NewPresignClient(s.Client)

	request, err := presignedClient.PresignGetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(objectKey),
		},
		s3.WithPresignExpires(15*time.Minute),
	)

	if err != nil {
		return "", err
	}

	return request.URL, nil
}
