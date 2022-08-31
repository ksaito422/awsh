package ecs

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type ECSDescribeTaskDefinitionAPI interface {
	DescribeTaskDefinition(ctx context.Context,
		params *ecs.DescribeTaskDefinitionInput,
		optFns ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error)
}

func describeTaskDefinitionAPI(c context.Context, api ECSDescribeTaskDefinitionAPI, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	return api.DescribeTaskDefinition(c, input)
}

// aws ecs describe-task-definition
func DescribeTask(cfg aws.Config, taskDef string) *types.TaskDefinition {
	client := ecs.NewFromConfig(cfg)
	input := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: &taskDef,
	}

	resp, err := describeTaskDefinitionAPI(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving list tasks:")
		fmt.Println(err)
		os.Exit(1)
	}

	return resp.TaskDefinition
}
