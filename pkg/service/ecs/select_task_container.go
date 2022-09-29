package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// Receives a value of type DescribeTasksOutput in the argument and returns container name in string.
func SelectTaskContainer(input *ecs.DescribeTasksOutput) string {
	// TODO: 引数のnullチェック入れる
	return *input.Tasks[0].Containers[0].Name
}
