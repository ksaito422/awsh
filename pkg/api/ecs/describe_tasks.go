package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ecs#Client.DescribeTasks
type ECSDescribeTasksAPI interface {
	DescribeTasks(ctx context.Context,
		params *ecs.DescribeTasksInput,
		optFns ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)
}

func describeTasksAPI(c context.Context, api ECSDescribeTasksAPI, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	return api.DescribeTasks(c, input)
}

// Returns detailed data for the selected ecs task.
// For aws cli -> aws ecs describe-tasks
func (s *ecsApi) DescribeTasks(cfg aws.Config, cluster, taskArn string) (*ecs.DescribeTasksOutput, error) {
	client := ecs.NewFromConfig(cfg)
	taskArr := []string{taskArn}
	input := &ecs.DescribeTasksInput{
		Cluster: &cluster,
		Tasks:   taskArr,
	}

	resp, err := describeTasksAPI(context.TODO(), client, input)
	if err != nil {
		return nil, errDescribeTask
	}

	return resp, nil
}
