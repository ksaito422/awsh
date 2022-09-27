package ecs

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ecsTasksName struct {
	List []string
}

func (m *ecsTasksName) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListTaskDefinitionsOutput in the argument and returns task definition in string.
func SelectTaskDefinition(input *ecs.ListTaskDefinitionsOutput) string {
	// TODO: 引数のnullチェック入れる
	ls := new(ecsTasksName)
	for _, task := range input.TaskDefinitionArns {
		ls.Set(task)
	}

	taskDef := prompt.ChooseValueFromPromptItems("Select ECS Task Definitions", ls.List)
	return taskDef
}
