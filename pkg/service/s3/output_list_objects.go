package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Output the bucket and object list information passed as arguments.
func OutputListObjects(v *s3.ListObjectsV2Output, bucket string) {
	// TODO: 引数のnullチェック入れる
	fmt.Println("Objects in " + bucket + ":")
	for _, item := range v.Contents {
		fmt.Println("Name:", *item.Key, " | ", "Last modified:", *item.LastModified, " | ", "Size:", item.Size, " | ", "Storage:", item.StorageClass)
	}
}
