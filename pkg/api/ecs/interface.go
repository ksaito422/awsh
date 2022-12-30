package ecs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type ECSApi interface {
	DescribeTaskDefinition(aws.Config, string) (*types.TaskDefinition, error)
	DescribeTasks(aws.Config, string, string) (*ecs.DescribeTasksOutput, error)
	ExecuteCommand(cfg aws.Config, cluster, taskArn, container, runtimeId string) error
	ListClusters(aws.Config) (*ecs.ListClustersOutput, error)
	ListTaskDefinitions(aws.Config) (*ecs.ListTaskDefinitionsOutput, error)
	ListTasks(aws.Config, string, string) (*ecs.ListTasksOutput, error)
}

type ecsApi struct{}

// constructor関数
func NewECSAPI() *ecsApi {
	return &ecsApi{}
}
