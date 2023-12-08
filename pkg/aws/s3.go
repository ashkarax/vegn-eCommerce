package aws

import (
	"fmt"
	"mime/multipart"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	uuidgenerator "github.com/ashkarax/vegn-eCommerce/pkg/uuid_generator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


type awsS3Service struct {
	s3Credentials config.AWS
}

var awss3service awsS3Service

func AWSS3ImageUploaderSetup(data config.AWS) {
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

func AWSImageUploader(file *multipart.FileHeader, sess *session.Session, bucketFolder *string) (*string, error) {
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
	})

	if err != nil {
		fmt.Println(err)
		return &upload.Location, err
	}

	return &upload.Location, nil

}
