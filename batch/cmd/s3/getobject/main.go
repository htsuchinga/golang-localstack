package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/htsuchinga/golang-localstack/batch/internal/awsapi"
	"github.com/htsuchinga/golang-localstack/batch/internal/logger"
	"github.com/htsuchinga/golang-localstack/config"
)

const batchName = "s3/getobject"

func main() {
	start := time.Now()
	logger.DefaultModuleName = batchName
	logger.Info("start %s", logger.Version())
	defer func() {
		err := recover()
		if err != nil {
			logger.Error("error:%v", err)
			os.Exit(1)
		}
	}()

	// S3Client
	s3Config := &awsapi.S3Config{
		EndPoint:   config.Params.S3Endpoint,
		BucketName: "test",
	}
	s3Client := awsapi.NewS3Client(s3Config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	_, err := s3Client.GetObject(ctx, "sample.csv")
	if err != nil {
		panic(fmt.Errorf("s3 get failed:%v", err))
	}

	sec := time.Since(start).Seconds()
	logger.Info("end (%.0f sec)", sec)
}
