package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3BucketsName struct {
	List []string
}

func (l *s3BucketsName) Set(v string) {
	l.List = append(l.List, v)
}

// Receives a value of type ListBucketsOutput in the argument and returns a list of bucket names in []string.
func CreateBucketsNameList(input *s3.ListBucketsOutput) []string {
	// TODO: 引数のnullチェック入れる
	ls := new(s3BucketsName)
	for _, bucket := range input.Buckets {
		ls.Set(*bucket.Name)
	}

	return ls.List
}
