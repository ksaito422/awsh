package s3

import (
	"fmt"

	"awsh/internal/logging"
	s3api "awsh/pkg/api/s3"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Output the bucket and object list information passed as arguments.
func ListObjects(cfg aws.Config) error {
	listBuckets, err := s3api.ListBuckets(cfg)
	if err != nil {
		return err
	}

	// バケットが一つもない場合
	if len(listBuckets.Buckets) == 0 {
		log := logging.Log()
		log.Info().Msg("No buckets")

		return nil
	}

	bucketName := SelectBucketName(listBuckets)
	listObjects, err := s3api.ListObjects(cfg, bucketName)
	if err != nil {
		return err
	}

	// オブジェクトが一つもない場合
	if len(listObjects.Contents) == 0 {
		log := logging.Log()
		log.Info().Msg("No objects")

		return nil
	}

	fmt.Println("Objects in " + bucketName + ":")
	for _, item := range listObjects.Contents {
		fmt.Println("Name:", *item.Key, " | ", "Last modified:", *item.LastModified, " | ", "Size:", item.Size, " | ", "Storage:", item.StorageClass)
	}

	return nil
}
