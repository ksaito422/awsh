//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/service/$GOFILE
package s3

import (
	s3 "awsh/pkg/api/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type S3Servicer interface {
	ListBuckets(aws.Config) error
	ListObjects(aws.Config) error
	DownloadObject(aws.Config) error
}

type S3Service struct {
	Api s3.S3Api
}

// constructor関数
func NewS3Service(s s3.S3Api) *S3Service {
	return &S3Service{Api: s}
}
