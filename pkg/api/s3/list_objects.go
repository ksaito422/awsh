package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3ListObjectsAPI interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

func getAllObjects(c context.Context, api s3ListObjectsAPI, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return api.ListObjectsV2(c, input)
}

// Returns data from the selected objects.
// For aws cli -> aws s3 list-object
func (s *s3Api) ListObjects(cfg aws.Config, bucket string) (*s3.ListObjectsV2Output, error) {
	client := s3.NewFromConfig(cfg)
	// 引数で指定したバケット内のオブジェクトの取得
	bucket_input := &s3.ListObjectsV2Input{
		Bucket: &bucket,
	}

	resp, err := getAllObjects(context.TODO(), client, bucket_input)
	if err != nil {
		return nil, errFetchObject
	}

	return resp, nil
}
