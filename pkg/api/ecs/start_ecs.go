package ecs

import (
	"context"
	"fmt"

	"awsh/internal/logging"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type ECSRunTaskAPI interface {
	RunTask(ctx context.Context,
		params *ecs.RunTaskInput,
		optFns ...func(*ecs.Options)) (*ecs.RunTaskOutput, error)
}

func myRunTask(c context.Context, api ECSRunTaskAPI, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	return api.RunTask(c, input)
}

/*
Starts a task based on the selected arguments.

For aws cli -> aws ecs run-task
*/
func (s *ecsApi) StartContainer(cfg aws.Config, cluster, taskArn, subnetId string) error {
	client := ecs.NewFromConfig(cfg)
	// TODO: セキュリティグループも指定する
	input := &ecs.RunTaskInput{
		TaskDefinition:       &taskArn,
		Cluster:              &cluster,
		EnableExecuteCommand: true,
		LaunchType:           types.LaunchTypeFargate,
		NetworkConfiguration: &types.NetworkConfiguration{
			AwsvpcConfiguration: &types.AwsVpcConfiguration{
				Subnets:        []string{subnetId},
				AssignPublicIp: "DISABLED",
			},
		},
	}

	_, err := myRunTask(context.TODO(), client, input)
	if err != nil {
		log := logging.Log()
		log.Error().Err(err).Msg("Got an error retrieving buckets:")

		return nil
	}

	fmt.Println("Done. Bye!")

	return nil
}
