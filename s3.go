package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s3Client)
	objInput := &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}

	presignedRequest, err := presignClient.PresignGetObject(context.Background(), objInput, s3.WithPresignExpires(expireTime))
	if err != nil {
		return "", err
	}
	return presignedRequest.URL, nil
}
