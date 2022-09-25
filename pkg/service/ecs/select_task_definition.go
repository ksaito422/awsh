package ecs

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (m *ECSTasksName) Set(value string) {
	m.List = append(m.List, value)
}

type ECSTasksName struct {
	List []string
}

func SelectTaskDefinition(input *ecs.ListTaskDefinitionsOutput) string {
	ls := new(ECSTasksName)
	for _, task := range input.TaskDefinitionArns {
		ls.Set(task)
	}

	taskDef := prompt.ChooseValueFromPromptItems("Select ECS Task Definitions", ls.List)
	return taskDef
}
