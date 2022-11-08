package service

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
)

type s3BucketsName struct {
	List []string
}

func (l *s3BucketsName) Set(v string) {
	l.List = append(l.List, v)
}

// Receives a value of type ListBucketsOutput in the argument and returns bucket name in string.
func SelectBucketName(input *s3.ListBucketsOutput) string {
	ls := new(s3BucketsName)
	for _, bucket := range input.Buckets {
		ls.Set(*bucket.Name)
	}

	bucket := prompt.ChooseValueFromPromptItems("Select S3 Buckets", ls.List)
	return bucket
}
