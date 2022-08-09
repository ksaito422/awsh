package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
)

type DownloadS3Client struct {
	downloader *manager.Downloader
	uploader   *manager.Uploader
	client     *s3.Client
}

// ファクトリー関数
func NewDownloadS3Client(cfg aws.Config) *DownloadS3Client {
	client := s3.NewFromConfig(cfg)
	downloader := manager.NewDownloader(client)
	uploader := manager.NewUploader(client)

	return &DownloadS3Client{
		downloader: downloader,
		uploader:   uploader,
		client:     client,
	}
}

// s3オブジェクトのダウンロード処理
func (c *DownloadS3Client) DownloadSingleObject(bucket, key, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	_, err := c.downloader.Download(context.Background(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("download successed")
}

// aws s3api get-object(download)
func DownloadObject(cfg aws.Config, bucket string, objects *s3.ListObjectsV2Output) {
	ss := new(S3ObjectsName)
	for _, item := range objects.Contents {
		ss.Set(*item.Key)
	}

	select_object := prompt.ChooseValueFromPromptItems("Select S3 Objects", ss.List)
	fmt.Println(select_object)

	client := NewDownloadS3Client(cfg)
	client.DownloadSingleObject(bucket, select_object, select_object)
}
