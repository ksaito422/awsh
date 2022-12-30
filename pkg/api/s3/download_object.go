package s3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
)

type S3ObjectsName struct {
	List []string
}

func (m *S3ObjectsName) Set(value string) {
	m.List = append(m.List, value)
}

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

// Download s3 object.
func (c *DownloadS3Client) DownloadSingleObject(bucket, key string) error {
	file, _ := os.Create(key)
	defer file.Close()

	_, err := c.downloader.Download(context.Background(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil
	}

	return nil
}

// Download selected objects.
func (s *s3Api) DownloadObject(cfg aws.Config, bucket string, objects *s3.ListObjectsV2Output) error {
	ss := new(S3ObjectsName)
	for _, item := range objects.Contents {
		ss.Set(*item.Key)
	}

	object := prompt.ChooseValueFromPromptItems("Select S3 Objects", ss.List)

	client := NewDownloadS3Client(cfg)
	if err := client.DownloadSingleObject(bucket, object); err != nil {
		return errDownloadObject
	}

	return nil
}
