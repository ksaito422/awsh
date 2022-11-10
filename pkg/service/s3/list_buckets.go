package s3

import (
	"fmt"

	"awsh/internal/logging"
	s3api "awsh/pkg/api/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// Outputs information about the bucket passed as argument.
func ListBuckets(cfg aws.Config) error {
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

	for _, bucket := range listBuckets.Buckets {
		fmt.Println(*bucket.Name)
	}

	return nil
}
