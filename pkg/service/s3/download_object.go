package s3

import (
	"fmt"

	"awsh/internal/logging"

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
		log := logging.Log()
		log.Info().Msg("No buckets")

		return nil
	}

	bucketName := s.Api.SelectBucketName(listBuckets)
	listObjects, err := s.Api.ListObjects(cfg, bucketName)
	if err != nil {
		return err
	}

	// オブジェクトが一つもない場合
	if len(listObjects.Contents) == 0 {
		log := logging.Log()
		log.Info().Msg("No objects")

		return nil
	}

	if err := s.Api.DownloadObject(cfg, bucketName, listObjects); err != nil {
		return err
	}

	fmt.Println("Download successed!")

	return nil
}
