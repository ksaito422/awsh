package ecs

import (
	"context"
	"fmt"
	"os"

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

// aws ecs run-task
func StartContainer(cfg aws.Config, cluster, taskArn, subnetId string) {
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
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done. Bye!")
}
