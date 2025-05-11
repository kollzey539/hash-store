package storage

import (
	"bytes"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Storage(accessKey, secretKey, region, bucket string) (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3Storage{
		Client:     client,
		BucketName: bucket,
	}, nil
}

func (s *S3Storage) PutItem(key, value string) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &s.BucketName,
		Key:         &key,
		Body:        bytes.NewReader([]byte(value)),
		ContentType: aws.String("text/plain"),
	})
	return err
}

func (s *S3Storage) GetItem(key string) (string, error) {
	output, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &s.BucketName,
		Key:    &key,
	})
	if err != nil {
		return "", err
	}
	defer output.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(output.Body)
	if err != nil {
		return "", err
	}

	if buf.Len() == 0 {
		return "", errors.New("empty object")
	}

	return buf.String(), nil
}
