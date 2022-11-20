package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Output the bucket and object list information passed as arguments.
func (s *S3Service) DownloadObject(cfg aws.Config) error {
	listBuckets, err := s.Api.ListBuckets(cfg)
	if err != nil {
		return err
	}

	// バケットが一つもない場合
	if len(listBuckets.Buckets) == 0 {
		return noBucket
	}

	bucketName := s.Api.SelectBucketName(listBuckets)
	listObjects, err := s.Api.ListObjects(cfg, bucketName)
	if err != nil {
		return err
	}

	// オブジェクトが一つもない場合
	if len(listObjects.Contents) == 0 {
		return noObject
	}

	if err := s.Api.DownloadObject(cfg, bucketName, listObjects); err != nil {
		return err
	}

	fmt.Println(OBJECT_DOWNLOAD)

	return nil
}
