package ecs

import (
	"context"
	"fmt"
	"os"

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
func ListTaskDefinitions(cfg aws.Config) *ecs.ListTaskDefinitionsOutput {
	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListTaskDefinitionsInput{}

	resp, err := GetAllTaskDefinitions(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving clusters:")
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}
