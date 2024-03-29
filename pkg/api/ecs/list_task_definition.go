package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSListTaskDefinitionsAPI interface {
	ListTaskDefinitions(ctx context.Context,
		params *ecs.ListTaskDefinitionsInput,
		optFns ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error)
}

func GetAllTaskDefinitions(c context.Context, api ECSListTaskDefinitionsAPI, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return api.ListTaskDefinitions(c, input)
}

// Returns data for the selected ecs task definition.
// For aws cli -> aws ecs list-task-definitions
func (s *ecsApi) ListTaskDefinitions(cfg aws.Config) (*ecs.ListTaskDefinitionsOutput, error) {
	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListTaskDefinitionsInput{}

	resp, err := GetAllTaskDefinitions(context.TODO(), client, input)
	if err != nil {
		return nil, errGetAllTask
	}

	return resp, nil
}
