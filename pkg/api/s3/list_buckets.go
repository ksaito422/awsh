package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func getAllBuckets(c context.Context, api s3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

// Returns data from the selected client's Bucket.
// For aws cli -> aws s3 ls
func ListBuckets(cfg aws.Config) *s3.ListBucketsOutput {
	client := s3.NewFromConfig(cfg)

	input := &s3.ListBucketsInput{}

	resp, err := getAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}
