//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/api/$GOFILE
package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Api interface {
	ListBuckets(aws.Config) (*s3.ListBucketsOutput, error)
	ListObjects(aws.Config, string) (*s3.ListObjectsV2Output, error)
	DownloadObject(aws.Config, string, *s3.ListObjectsV2Output) error
	SelectBucketName(*s3.ListBucketsOutput) string
}

type s3Api struct{}

// constructor関数
func NewS3API() *s3Api {
	return &s3Api{}
}
