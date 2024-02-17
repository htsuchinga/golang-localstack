package awsapi

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	S3Client interface {
		PutObject(key string, body io.ReadSeeker) (*s3.PutObjectOutput, error)
		GetObject(ctx context.Context, key string) (*s3.GetObjectOutput, error)
	}

	S3Config struct {
		EndPoint   string
		BucketName string
	}

	s3Client struct {
		clientCfg *S3Config
		sdkCfg    aws.Config
	}
)

func NewS3Client(cfg *S3Config) S3Client {
	loadOptions := []func(*config.LoadOptions) error{config.WithRegion("ap-northeast-1")}

	if cfg.EndPoint != "" {
		endpoint := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: cfg.EndPoint,
			}, nil
		})
		loadOptions = append(loadOptions, config.WithEndpointResolverWithOptions(endpoint))
	}

	sdkCfg, err := config.LoadDefaultConfig(context.TODO(), loadOptions...)
	if err != nil {
		panic(err)
	}
	return &s3Client{
		clientCfg: cfg,
		sdkCfg:    sdkCfg,
	}
}

func (c *s3Client) PutObject(key string, body io.ReadSeeker) (*s3.PutObjectOutput, error) {
	api := s3.NewFromConfig(c.sdkCfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	return api.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.clientCfg.BucketName),
		Key:    aws.String(key),
		Body:   body,
	})
}

func (c *s3Client) GetObject(ctx context.Context, key string) (*s3.GetObjectOutput, error) {
	api := s3.NewFromConfig(c.sdkCfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})

	return api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.clientCfg.BucketName),
		Key:    aws.String(key),
	})
}
