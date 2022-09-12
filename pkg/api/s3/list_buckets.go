package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

/*
Returns data from the selected bucket.

For aws cli -> aws s3 ls
*/
func ListBuckets(cfg aws.Config) []string {
	client := s3.NewFromConfig(cfg)

	input := &s3.ListBucketsInput{}

	resp, err := GetAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		os.Exit(1)
	}

	ss := new(S3BucketsName)
	for _, bucket := range resp.Buckets {
		ss.Set(*bucket.Name)
	}

	return ss.List
}
