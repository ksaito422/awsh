package pkg

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ListObjectsAPI interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

func GetObjects(c context.Context, api S3ListObjectsAPI, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return api.ListObjectsV2(c, input)
}

func S3ListObjects (cfg aws.Config, select_bucket string) {
	client := s3.NewFromConfig(cfg)
	// 上で選択したバケット内のオブジェクトの取得
	bucket_input := &s3.ListObjectsV2Input{
		Bucket: &select_bucket,
	}

	resp, err := GetObjects(context.TODO(), client, bucket_input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		return
	}

	fmt.Println("Objects in " + select_bucket + ":")
	for _, item := range resp.Contents {
		fmt.Println("Name:", *item.Key, " | ", "Last modified:", *item.LastModified, " | ", "Size:", item.Size, " | ", "Storage:", item.StorageClass)
	}

}

