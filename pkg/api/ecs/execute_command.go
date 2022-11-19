package ecs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type ECSExecuteCommandAPI interface {
	ExecuteCommand(ctx context.Context,
		params *ecs.ExecuteCommandInput,
		optFns ...func(*ecs.Options)) (*ecs.ExecuteCommandOutput, error)
}

func executeCommandAPI(c context.Context, api ECSExecuteCommandAPI, input *ecs.ExecuteCommandInput) (*ecs.ExecuteCommandOutput, error) {
	return api.ExecuteCommand(c, input)
}

/*
Connect to the selected container with ecs-exec.

For aws cli -> aws ecs execute-command
*/
func (s *ecsApi) ExecuteCommand(cfg aws.Config, cluster, taskArn, container, runtimeId string) error {
	client := ecs.NewFromConfig(cfg)
	sh := "/bin/sh"
	input := &ecs.ExecuteCommandInput{
		Cluster:     &cluster,
		Task:        &taskArn,
		Container:   &container,
		Interactive: true,
		Command:     &sh,
	}

	resp, err := executeCommandAPI(context.TODO(), client, input)
	if err != nil {
		return errExecuteCommand
	}

	sess, _ := json.Marshal(resp.Session)
	target := fmt.Sprintf("ecs:%s_%s_%s", cluster, taskArn, runtimeId)
	ssmTarget := ssm.StartSessionInput{
		Target: &target,
	}
	targetJSON, err := json.Marshal(ssmTarget)

	// TODO: aws sdk ssmの方でできないのか？
	// StartSession実行時の引数にJSON化したExecuteCommandの返り値を与える
	cmd := exec.Command(
		"session-manager-plugin",
		string(sess),
		"ap-northeast-1",
		"StartSession",
		"",
		string(targetJSON),
		"https://ssm.ap-northeast-1.amazonaws.com",
	)
	signal.Ignore(os.Interrupt)
	defer signal.Reset(os.Interrupt)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
