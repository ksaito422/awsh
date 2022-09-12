package ecs

import (
	"context"
	"fmt"
	"os"

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

/*
Returns detailed data for the selected ecs task.

For aws cli -> aws ecs describe-tasks
*/
func DescribeTasks(cfg aws.Config, cluster, taskArn string) (string, string) {
	client := ecs.NewFromConfig(cfg)
	taskArr := []string{taskArn}
	input := &ecs.DescribeTasksInput{
		Cluster: &cluster,
		Tasks:   taskArr,
	}

	resp, err := describeTasksAPI(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving describe tasks:")
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: 戻り値の設計をして、returnを変えたい
	return *resp.Tasks[0].Containers[0].Name, *resp.Tasks[0].Containers[0].RuntimeId
}
