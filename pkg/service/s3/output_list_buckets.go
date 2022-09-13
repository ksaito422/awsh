package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Outputs information about the bucket passed as argument.
func OutputListBuckets(value *s3.ListBucketsOutput) {
	// TODO: 引数のnullチェック入れる
	for _, bucket := range value.Buckets {
		fmt.Println(*bucket.Name)
	}
}
