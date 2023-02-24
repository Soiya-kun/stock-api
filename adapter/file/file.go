package file

import (
	"fmt"
	"time"

	"gitlab.com/soy-app/stock-api/config"
	"gitlab.com/soy-app/stock-api/usecase/port"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	awsDriver "gitlab.com/soy-app/stock-api/adapter/aws"
)

type File struct {
	awsCli *awsDriver.Cli
}

func NewFileDriver(awsCli *awsDriver.Cli) port.FileDriver {
	return &File{
		awsCli: awsCli,
	}
}

func (f File) CopyFile(srcId, dstId string) error {
	sess, err := f.awsCli.CreateSession()
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}
	svc := s3.New(sess)
	bucket := aws.String(config.S3Bucket())
	copySource := aws.String(config.S3Bucket() + "/" + srcId)
	key := aws.String(dstId)

	if _, err := svc.CopyObject(&s3.CopyObjectInput{
		CopySource: copySource,
		Bucket:     bucket,
		Key:        key,
	}); err != nil {
		return fmt.Errorf("copy object: %w", err)
	}

	return nil
}

func (f File) CreatePreSignedURLForGet(filepath string) (string, error) {
	sess, err := f.awsCli.CreateSession()
	if err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}

	svc := s3.New(sess)
	key := aws.String(filepath)
	bucket := aws.String(config.S3Bucket())

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	url, err := req.Presign(5 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("create pre signed url for get: %w", err)
	}

	return url, nil
}

func (f File) CreatePreSignedURLForPut(filepath string) (string, error) {
	sess, err := f.awsCli.CreateSession()
	if err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}

	svc := s3.New(sess)
	key := aws.String(filepath)
	bucket := aws.String(config.S3Bucket())

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	url, err := req.Presign(5 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("create pre signed url for put: %w", err)
	}

	return url, nil
}

func (f File) DeleteFileWithPath(filepath string) error {
	sess, err := f.awsCli.CreateSession()
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}

	svc := s3.New(sess)
	key := aws.String(filepath)
	bucket := aws.String(config.S3Bucket())

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: bucket,
		Key:    key,
	})

	if err != nil {
		return err
	}

	return nil
}
