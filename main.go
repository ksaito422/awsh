package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/manifoldco/promptui"
)

type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func chooseValueFromPrompt(l string, d string) string {
	prompt := promptui.Prompt{
		Label:   l,
		Default: d,
	}
	v, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return v
}

func chooseValueFromPromptItems(l string, i []string) string {
	prompt := promptui.Select{
		Label: l,
		Items: i,
	}
	_, v, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return v
}

func GetAllBuckets(c context.Context, api S3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

func main() {
	aws_region := chooseValueFromPrompt("Please enter aws region(Default: ap-northeast-1)", "ap-northeast-1")
	aws_profile := chooseValueFromPrompt("Please enter aws profile(If empty, default settings are loaded)", "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(aws_region), config.WithSharedConfigProfile(aws_profile))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.ListBucketsInput{}

	result, err := GetAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		return
	}

	fmt.Println("Buckets:")

	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}
}

