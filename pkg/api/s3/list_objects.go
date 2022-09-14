package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
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
func ListObjects(cfg aws.Config, listBuckets []string) (*s3.ListObjectsV2Output, string) {
	bucket := prompt.ChooseValueFromPromptItems("Select S3 Buckets", listBuckets)

	client := s3.NewFromConfig(cfg)
	// 上で選択したバケット内のオブジェクトの取得
	bucket_input := &s3.ListObjectsV2Input{
		Bucket: &bucket,
	}

	resp, err := getAllObjects(context.TODO(), client, bucket_input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		os.Exit(1)
	}

	return resp, bucket
}
