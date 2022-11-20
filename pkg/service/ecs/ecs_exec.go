package ecs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func (s *ECSService) EcsExec(cfg aws.Config) error {
	listClusters, err := s.Api.ListClusters(cfg)
	if err != nil {
		return err
	}
	// ECS clusterが一つもない場合
	if len(listClusters.ClusterArns) == 0 {
		return noEcsCluster
	}

	clusterArn := SelectClusterArn(listClusters)
	listTaskDefs, err := s.Api.ListTaskDefinitions(cfg)
	if err != nil {
		return err
	}
	// Task definitionが一つもない場合
	if len(listTaskDefs.TaskDefinitionArns) == 0 {
		return noTaskDefinition
	}

	taskDef := SelectTaskDefinition(listTaskDefs)
	taskDefDetail, err := s.Api.DescribeTaskDefinition(cfg, taskDef)
	if err != nil {
		return err
	}

	listTasks, err := s.Api.ListTasks(cfg, clusterArn, *taskDefDetail.Family)
	if err != nil {
		return err
	}
	// task listが一つもない場合
	if len(listTasks.TaskArns) == 0 {
		return noEcsTask
	}

	taskArn := SelectTaskArn(listTasks)
	taskDetail, err := s.Api.DescribeTasks(cfg, clusterArn, taskArn)
	if err != nil {
		return nil
	}

	containerName := SelectTaskContainer(taskDetail)
	runtimeId := SelectRuntimeId(taskDetail)
	if err := s.Api.ExecuteCommand(cfg, clusterArn, taskArn, containerName, runtimeId); err != nil {
		return err
	}

	return nil
}
