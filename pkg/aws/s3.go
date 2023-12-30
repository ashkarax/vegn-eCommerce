package aws

import (
	"bytes"
	"fmt"
	"mime/multipart"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	uuidgenerator "github.com/ashkarax/vegn-eCommerce/pkg/uuid_generator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type awsS3Service struct {
	s3Credentials config.AWS
}

var awss3service awsS3Service

func AWSS3FileUploaderSetup(data config.AWS) {
	awss3service.s3Credentials = data
}

func AWSSessionInitializer() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awss3service.s3Credentials.Region),
			Credentials: credentials.NewStaticCredentials(
				awss3service.s3Credentials.AccessKey,
				awss3service.s3Credentials.SecrectKey,
				"",
			),
			Endpoint: aws.String(awss3service.s3Credentials.Endpoint),
		},
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func AWSS3ImageUploader(file *multipart.FileHeader, sess *session.Session, bucketFolder *string) (*string, error) {
	var nullstringresponse string

	image, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return &nullstringresponse, err
	}
	defer image.Close()

	randomName := uuidgenerator.ReturnUuid()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucketFolder),
		Key:    aws.String(*randomName),
		Body:   image,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		fmt.Println(err)
		return &upload.Location, err
	}

	return &upload.Location, nil

}

func AWSS3ByteBufferUploader(content []byte, sess *session.Session, bucketFolder *string) (*string, error) {

	randomName := uuid.New().String()


	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucketFolder),
		Key:    aws.String(randomName),
		Body:   bytes.NewReader(content),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return &upload.Location, err
	}

	return &upload.Location, nil
}

func AWSS3XLSXUploader(content []byte, sess *session.Session, bucketFolder *string) (*string, error) {

	randomName := uuid.New().String()

	verr:=fmt.Sprintf("%s.xlsx",randomName)

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucketFolder),
		Key:    aws.String(verr),
		Body:   bytes.NewReader(content),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return &upload.Location, err
	}

	return &upload.Location, nil
}
