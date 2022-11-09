package ecs

import (
	"awsh/internal/logging"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSListTasksAPI interface {
	ListTasks(ctx context.Context,
		params *ecs.ListTasksInput,
		optFns ...func(*ecs.Options)) (*ecs.ListTasksOutput, error)
}

type ECSListTasks struct {
	List []string
}

func (m *ECSListTasks) Set(value string) {
	m.List = append(m.List, value)
}

func listTaskAPI(c context.Context, api ECSListTasksAPI, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	return api.ListTasks(c, input)
}

// Returns data for the selected ecs task.
// For aws cli -> aws ecs list-tasks
func ListTasks(cfg aws.Config, cluster, family string) (*ecs.ListTasksOutput, error) {
	client := ecs.NewFromConfig(cfg)
	input := &ecs.ListTasksInput{
		Cluster: &cluster,
		Family:  &family,
	}

	resp, err := listTaskAPI(context.TODO(), client, input)
	if err != nil {
		log := logging.Log()
		log.Error().Err(err).Msg("Got an error retrieving list tasks:")

		return nil, err
	}

	return resp, nil
}
