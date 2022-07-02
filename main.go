package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/manifoldco/promptui"

	"awsh/pkg"
)

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

func main() {
	aws_region := chooseValueFromPrompt("Please enter aws region(Default: ap-northeast-1)", "ap-northeast-1")
	aws_profile := chooseValueFromPrompt("Please enter aws profile(If empty, default settings are loaded)", "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(aws_region), config.WithSharedConfigProfile(aws_profile))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	select_bucket := pkg.S3ListBuckets(cfg)
	pkg.S3ListObjects(cfg, select_bucket)
}
