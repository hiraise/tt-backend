package s3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Service struct {
	bucket string
	client *s3.Client
}

func New(accessKey, secretKey, endpoint, bucket string) (*Service, error) {

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			),
		),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, fmt.Errorf("loading s3 configuration failed: %w", err)
	}
	retVal := &Service{
		bucket: bucket,
	}
	retVal.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
	return retVal, nil
}

func (s *Service) Save(ctx context.Context, file []byte, name string, mimeType string) error {

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(name),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(mimeType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return fmt.Errorf("cant upload file: %w", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, name string) error {
	_, err := s.client.DeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(name),
		},
	)
	if err != nil {
		return fmt.Errorf("cant delete file: %w", err)
	}
	return nil
}

func (s *Service) GetPath(name string) string {
	return fmt.Sprintf("%s/%s/%s", *s.client.Options().BaseEndpoint, s.bucket, name)
}
