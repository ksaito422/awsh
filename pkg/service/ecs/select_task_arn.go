package ecs

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ecsTaskArn struct {
	List []string
}

func (m *ecsTaskArn) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListTasksOutput in the argument and returns task arn in string.
func SelectTaskArn(input *ecs.ListTasksOutput) string {
	// TODO: 引数のnullチェック入れる
	ls := new(ecsTaskArn)
	for _, task := range input.TaskArns {
		ls.Set(task)
	}

	taskArn := prompt.ChooseValueFromPromptItems("Select ECS Task", ls.List)
	return taskArn
}
