package endpoints

import (
	// "awsh/pkg/api/ec2"
	// ecsapi "awsh/pkg/api/ecs"
	// ecsservice "awsh/pkg/service/ecs"
	"awsh/pkg/service/s3"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Routing of operation actions on AWS resources.
func Controller(cfg aws.Config, action string) error {
	switch action {
	// S3
	case "ListBuckets":
		err := s3.ListBuckets(cfg)
		return err

	case "ListObjects":
		err := s3.ListObjects(cfg)
		return err

	case "DownloadObject":
		err := s3.DownloadObject(cfg)
		return err

		// ECS
		// case "StartECS":
		// 	// TODO: リファクタする
		// 	subnetId := ec2.DescribeSubnets(cfg)
		// 	// TODO: リファクタする
		// 	ec2.DescribeSecurityGroups(cfg)
		// 	listClusters := ecsapi.ListClusters(cfg)
		// 	clusterArn := ecsservice.SelectClusterArn(listClusters)
		// 	listTaskDefs := ecsapi.ListTaskDefinitions(cfg)
		// 	taskDef := ecsservice.SelectTaskDefinition(listTaskDefs)
		// 	taskDefDetail := ecsapi.DescribeTaskDefinition(cfg, taskDef)
		// 	// TODO: 起動するタスクにアタッチするセキュリティグループを後で渡す
		// 	ecsapi.StartContainer(cfg, clusterArn, *taskDefDetail.TaskDefinitionArn, *subnetId)
		//
		// case "ecs-exec":
		// 	listClusters := ecsapi.ListClusters(cfg)
		// 	clusterArn := ecsservice.SelectClusterArn(listClusters)
		// 	listTaskDefs := ecsapi.ListTaskDefinitions(cfg)
		// 	taskDef := ecsservice.SelectTaskDefinition(listTaskDefs)
		// 	taskDefDetail := ecsapi.DescribeTaskDefinition(cfg, taskDef)
		// 	listTasks := ecsapi.ListTasks(cfg, clusterArn, *taskDefDetail.Family)
		// 	taskArn := ecsservice.SelectTaskArn(listTasks)
		// 	taskDetail := ecsapi.DescribeTasks(cfg, clusterArn, taskArn)
		// 	containerName := ecsservice.SelectTaskContainer(taskDetail)
		// 	runtimeId := ecsservice.SelectRuntimeId(taskDetail)
		// 	ecsapi.ExecuteCommand(cfg, clusterArn, taskArn, containerName, runtimeId)
		//
		// case "StopECSTask":
		// 	listClusters := ecsapi.ListClusters(cfg)
		// 	clusterArn := ecsservice.SelectClusterArn(listClusters)
		// 	listTaskDefs := ecsapi.ListTaskDefinitions(cfg)
		// 	taskDef := ecsservice.SelectTaskDefinition(listTaskDefs)
		// 	taskDefDetail := ecsapi.DescribeTaskDefinition(cfg, taskDef)
		// 	listTasks := ecsapi.ListTasks(cfg, clusterArn, *taskDefDetail.Family)
		// 	taskArn := ecsservice.SelectTaskArn(listTasks)
		// 	ecsapi.StopTask(cfg, clusterArn, taskArn)
	default:
		return nil
	}
}
