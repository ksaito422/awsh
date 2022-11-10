package ecs

import (
	"awsh/internal/logging"
	ec2api "awsh/pkg/api/ec2"
	ecsapi "awsh/pkg/api/ecs"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func StartEcs(cfg aws.Config) error {
	// TODO: リファクタする
	subnetId, err := ec2api.DescribeSubnets(cfg)
	if err != nil {
		return err
	}
	// TODO: リファクタする
	if err := ec2api.DescribeSecurityGroups(cfg); err != nil {
		return err
	}

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

	// TODO: 起動するタスクにアタッチするセキュリティグループを後で渡す
	if err := ecsapi.StartContainer(cfg, clusterArn, *taskDefDetail.TaskDefinitionArn, *subnetId); err != nil {
		return err
	}

	return nil
}
