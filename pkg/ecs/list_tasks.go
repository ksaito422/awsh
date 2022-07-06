package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"

	"awsh/pkg/prompt"
)

type ECSListTaskDefinitionsAPI interface {
	ListTaskDefinitions(ctx context.Context,
		params *ecs.ListTaskDefinitionsInput,
		optFns ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error)
}

type ECSTasksName struct {
	List []string
}

func (m *ECSTasksName) Set(value string) {
	m.List = append(m.List, value)
}

func GetAllTaskDefinitions(c context.Context, api ECSListTaskDefinitionsAPI, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return api.ListTaskDefinitions(c, input)
}

// aws ecs list-task-definitions
func ListTaskDefinitions(cfg aws.Config) string {
	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListTaskDefinitionsInput{}

	result, err := GetAllTaskDefinitions(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving clusters:")
		fmt.Println(err)
		return "error"
	}

	ss := new(ECSTasksName)
	for _, task := range result.TaskDefinitionArns {
		ss.Set(task)
	}

	select_task_definition := prompt.ChooseValueFromPromptItems("Select ECS Task Definitions", ss.List)

	return select_task_definition
}

