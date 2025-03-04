package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Upload(fp string, Name string, region string) (string, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return "", err
	}

	file, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()

	u := s3manager.NewUploader(session)
	result, err := u.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Name),
		Key:    aws.String(fp),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
