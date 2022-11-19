package ecs

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"golang.org/x/xerrors"
)

var (
	noEcsCluster     = xerrors.New("ECS cluster does not exist.")
	noEcsTask        = xerrors.New("ECS task does not exist.")
	noTaskDefinition = xerrors.New("Task definition does not exist.")
)

type ecsClustersName struct {
	List []string
}

func (m *ecsClustersName) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListClustersOutput in the argument and returns cluster arn in string.
func SelectClusterArn(input *ecs.ListClustersOutput) string {
	ls := new(ecsClustersName)
	for _, cluster := range input.ClusterArns {
		ls.Set(cluster)
	}

	cluster := prompt.ChooseValueFromPromptItems("Select ECS Clusters", ls.List)
	return cluster
}

type ecsTasksName struct {
	List []string
}

func (m *ecsTasksName) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListTaskDefinitionsOutput in the argument and returns task definition in string.
func SelectTaskDefinition(input *ecs.ListTaskDefinitionsOutput) string {
	ls := new(ecsTasksName)
	for _, task := range input.TaskDefinitionArns {
		ls.Set(task)
	}

	taskDef := prompt.ChooseValueFromPromptItems("Select ECS Task Definitions", ls.List)
	return taskDef
}

type ecsTaskArn struct {
	List []string
}

func (m *ecsTaskArn) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListTasksOutput in the argument and returns task arn in string.
func SelectTaskArn(input *ecs.ListTasksOutput) string {
	ls := new(ecsTaskArn)
	for _, task := range input.TaskArns {
		ls.Set(task)
	}

	taskArn := prompt.ChooseValueFromPromptItems("Select ECS Task", ls.List)
	return taskArn
}

// Receives a value of type DescribeTasksOutput in the argument and returns container name in string.
func SelectTaskContainer(input *ecs.DescribeTasksOutput) string {
	// TODO: 引数のnullチェック入れる
	return *input.Tasks[0].Containers[0].Name
}

// Receives a value of type DescribeTasksOutput in the argument and returns runtime id in string.
func SelectRuntimeId(input *ecs.DescribeTasksOutput) string {
	return *input.Tasks[0].Containers[0].RuntimeId
}
