package s3

import (
	"fmt"

	"awsh/internal/logging"

	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/xerrors"
)

// Outputs information about the bucket passed as argument.
func (s *S3Service) ListBuckets(cfg aws.Config) error {
	listBuckets, err := s.Api.ListBuckets(cfg)
	if err != nil {
		return err
	}

	// バケットが一つもない場合
	if len(listBuckets.Buckets) == 0 {
		log := logging.Log()
		log.Info().Msg("No buckets")

		return xerrors.Errorf("No buckets")
	}

	for _, bucket := range listBuckets.Buckets {
		fmt.Println(*bucket.Name)
	}

	return nil
}
