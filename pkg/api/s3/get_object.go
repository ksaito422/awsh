package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"awsh/pkg/prompt"
)

type S3ObjectsName struct {
	List []string
}

func (m *S3ObjectsName) Set(value string) {
	m.List = append(m.List, value)
}

type S3GetObjectAPI interface {
	GetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

func GetObjects(c context.Context, api S3GetObjectAPI, input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return api.GetObject(c, input)
}

/*
Returns data from the selected object.

For aws cli -> aws s3api get-object
*/
func GetObject(cfg aws.Config, bucket string, objects *s3.ListObjectsV2Output) *s3.GetObjectOutput {
	ss := new(S3ObjectsName)
	for _, item := range objects.Contents {
		ss.Set(*item.Key)
	}

	select_object := prompt.ChooseValueFromPromptItems("Select S3 Objects", ss.List)
	fmt.Println(select_object)

	client := s3.NewFromConfig(cfg)
	object_input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &select_object,
	}

	resp, err := GetObjects(context.TODO(), client, object_input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}
