package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSStopTaskAPI interface {
	StopTask(ctx context.Context,
		params *ecs.StopTaskInput,
		optFns ...func(*ecs.Options)) (*ecs.StopTaskOutput, error)
}

func stopTaskAPI(c context.Context, api ECSStopTaskAPI, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	return api.StopTask(c, input)
}

/*
Stops the tasks of the selected cluster.

For aws cli -> aws ecs stop-task
*/
func (s *ecsApi) StopTask(cfg aws.Config, cluster, taskArn string) error {
	client := ecs.NewFromConfig(cfg)
	input := &ecs.StopTaskInput{
		Cluster: &cluster,
		Task:    &taskArn,
	}

	_, err := stopTaskAPI(context.TODO(), client, input)
	if err != nil {
		return errStopTask
	}

	fmt.Println("Done. Bye!")

	return nil
}
