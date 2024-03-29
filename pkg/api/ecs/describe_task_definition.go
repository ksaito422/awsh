package ecs

import (
	"context"

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

// Returns data for the ecs task definition.
// For aws cli -> aws ecs describe-task-definition
func (s *ecsApi) DescribeTaskDefinition(cfg aws.Config, taskDef string) (*types.TaskDefinition, error) {
	client := ecs.NewFromConfig(cfg)
	input := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: &taskDef,
	}

	resp, err := describeTaskDefinitionAPI(context.TODO(), client, input)
	if err != nil {
		return nil, errDescribeTaskDefinition
	}

	return resp.TaskDefinition, nil
}
