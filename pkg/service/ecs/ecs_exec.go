package ecs

import (
	"awsh/internal/logging"
	ecsapi "awsh/pkg/api/ecs"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func (s *ECS) EcsExec(cfg aws.Config) error {
	listClusters, err := ecsapi.ListClusters(cfg)
	if err != nil {
		return err
	}
	// ECS clusterが一つもない場合
	if len(listClusters.ClusterArns) == 0 {
		log := logging.Log()
		log.Info().Msg("Doesn't exist ECS cluster")

		return nil
	}

	clusterArn := SelectClusterArn(listClusters)
	listTaskDefs, err := ecsapi.ListTaskDefinitions(cfg)
	if err != nil {
		return err
	}
	// Task definitionが一つもない場合
	if len(listTaskDefs.TaskDefinitionArns) == 0 {
		log := logging.Log()
		log.Info().Msg("Doesn't exist ECS task definition")

		return nil
	}

	taskDef := SelectTaskDefinition(listTaskDefs)
	taskDefDetail, err := ecsapi.DescribeTaskDefinition(cfg, taskDef)
	if err != nil {
		return err
	}

	listTasks, err := ecsapi.ListTasks(cfg, clusterArn, *taskDefDetail.Family)
	if err != nil {
		return err
	}
	// task listが一つもない場合
	if len(listTasks.TaskArns) == 0 {
		log := logging.Log()
		log.Info().Msg("Doesn't exist ECS task")

		return nil
	}

	taskArn := SelectTaskArn(listTasks)
	taskDetail, err := ecsapi.DescribeTasks(cfg, clusterArn, taskArn)
	if err != nil {
		return nil
	}

	containerName := SelectTaskContainer(taskDetail)
	runtimeId := SelectRuntimeId(taskDetail)
	if err := ecsapi.ExecuteCommand(cfg, clusterArn, taskArn, containerName, runtimeId); err != nil {
		return err
	}

	return nil
}
