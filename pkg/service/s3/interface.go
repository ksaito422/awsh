package s3

import "github.com/aws/aws-sdk-go-v2/aws"

type S3Servicer interface {
	ListBuckets(aws.Config) error
	ListObjects(aws.Config) error
	DownloadObject(aws.Config) error
}

type S3 struct {
	service S3Servicer
}

// constructor関数
func NewS3Service(s S3Servicer) *S3 {
	return &S3{service: s}
}
