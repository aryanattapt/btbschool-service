package service

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateSessionS3() (sess *session.Session, err error) {
	sess, err = session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	return
}

func CreateBucketS3(client *s3.S3, bucketName string) (err error) {
	_, err = client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func ListBucketsS3(client *s3.S3) (res *s3.ListBucketsOutput, err error) {
	res, err = client.ListBuckets(nil)
	return
}

func UploadFileS3(uploader *s3manager.Uploader, file multipart.File, bucketName string, fileName string) (res *s3manager.UploadOutput, err error) {
	res, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	return
}
