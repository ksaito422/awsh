package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2DescribeRegionsAPI interface {
	DescribeRegions(ctx context.Context,
		params *ec2.DescribeRegionsInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
}

type regionName struct {
	List []string
}

func (m *regionName) Set(value string) {
	m.List = append(m.List, value)
}

// AWSがサポートしているRegionを取得する [aws ec2 describe-regions]
func GetAllRegions(c context.Context, api EC2DescribeRegionsAPI, input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	return api.DescribeRegions(c, input)
}

func Cfg() aws.Config {
	aws_profile := prompt.ChooseValueFromPrompt("Please enter aws profile(If empty, default settings are loaded)", "")

	// 仮でIAMロールの認証情報を取得。ec2 describe-regionsで利用可能なリージョンを取得するため
	ec2Cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	ec2Client := ec2.NewFromConfig(ec2Cfg)
	input := &ec2.DescribeRegionsInput{}
	resp, err := GetAllRegions(context.TODO(), ec2Client, input)
	if err != nil {
		fmt.Println("Got an error retrieving regions")
		fmt.Println(err)
		os.Exit(1)
	}

	// region nameのslice作成
	ss := new(regionName)
	for _, region := range (*resp).Regions {
		ss.Set(*region.RegionName)
	}

	// よく利用するリージョンを選択肢にする
	aws_region := prompt.ChooseValueFromPromptItems("Select aws region", []string{
		"ap-northeast-1",
		"ap-northeast-3",
		"us-east-1",
		"other",
	})
	// aws ec2 describe-regionsで返されたリージョンを選択肢にする
	if aws_region == "other" {
		aws_region = prompt.ChooseValueFromPromptItems("Select aws region", ss.List)
	}

	// IAMロールの認証情報で取得
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(aws_region), config.WithSharedConfigProfile(aws_profile))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
