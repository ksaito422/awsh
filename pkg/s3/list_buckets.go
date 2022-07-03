package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
)

type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

type S3BucketsName struct {
	List []string
}

func (m *S3BucketsName) Set(value string) {
	m.List = append(m.List, value)
}

func GetAllBuckets(c context.Context, api S3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

// aws s3 ls
func ListBuckets(cfg aws.Config) string {
	client := s3.NewFromConfig(cfg)

	input := &s3.ListBucketsInput{}

	result, err := GetAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		return "error"
	}

	ss := new(S3BucketsName)
	for _, bucket := range result.Buckets {
		ss.Set(*bucket.Name)
	}

	select_bucket := prompt.ChooseValueFromPromptItems("Select S3 Buckets", ss.List)

	return select_bucket
}
