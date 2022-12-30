package s3

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/xerrors"
)

var (
	errFetchBucket    = xerrors.New("An error occurred while retrieving the bucket.")
	errFetchObject    = xerrors.New("An error occurred while retrieving the object.")
	errDownloadObject = xerrors.New("An error occurred while downloading the object.")
)

type s3BucketsName struct {
	List []string
}

func (l *s3BucketsName) Set(v string) {
	l.List = append(l.List, v)
}

// Receives a value of type ListBucketsOutput in the argument and returns bucket name in string.
func (s *s3Api) SelectBucketName(input *s3.ListBucketsOutput) string {
	ls := new(s3BucketsName)
	for _, bucket := range input.Buckets {
		ls.Set(*bucket.Name)
	}

	bucket := prompt.ChooseValueFromPromptItems("Select S3 Buckets", ls.List)
	return bucket
}
