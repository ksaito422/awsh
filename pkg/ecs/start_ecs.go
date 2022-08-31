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
func StartContainer(cfg aws.Config) *ecs.RunTaskOutput {
	client := ecs.NewFromConfig(cfg)
	// TODO: inputに必要なものを与える
	input := &ecs.RunTaskInput{
		TaskDefinition:       "a", // 登録しているタスク定義を選択する
		Cluster:              "b", // 作成済みのクラスターを選択する
		EnableExecuteCommand: true,
		LaunchType:           types.LaunchTypeFargate,
		NetworkConfiguration: "c", // subnet、sgを選択したネットワーク設定を選択する
	}

	resp, err := myRunTask(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp)
	return resp
}

// aws ecs list-task-definitions --status active
// task-definitions ARN or リビジョン

// aws ecs list-clusters
// ARN or クラスター名
