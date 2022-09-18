package endpoints

import (
	"awsh/pkg/api/ec2"
	"awsh/pkg/api/ecs"
	s3api "awsh/pkg/api/s3"
	s3service "awsh/pkg/service/s3"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Routing of operation actions on AWS resources.
func Controller(cfg aws.Config, action string) {
	switch action {
	// S3
	case "ListBuckets":
		listBucketsOutput := s3api.ListBuckets(cfg)
		s3service.OutputListBuckets(listBucketsOutput)

	case "ListObjects":
		listBuckets := s3api.ListBuckets(cfg)
		listBucketsName := s3service.CreateBucketsNameList(listBuckets)
		listObjects, bucket := s3api.ListObjects(cfg, listBucketsName)
		s3service.OutputListObjects(listObjects, bucket)

	case "DownloadObject":
		listBuckets := s3api.ListBuckets(cfg)
		listBucketsName := s3service.CreateBucketsNameList(listBuckets)
		listObjects, bucket := s3api.ListObjects(cfg, listBucketsName)
		s3api.DownloadObject(cfg, bucket, listObjects)

	// ECS
	case "StartECS":
		subnetId := ec2.DescribeSubnets(cfg)
		// ec2.DescribeSecurityGroups(cfg)
		cluster := ecs.ListClusters(cfg)
		taskDef := ecs.ListTaskDefinitions(cfg)
		taskDefDetail := ecs.DescribeTaskDefinition(cfg, taskDef)
		// TODO: 起動するタスクにアタッチするセキュリティグループを後で渡す
		ecs.StartContainer(cfg, cluster, *taskDefDetail.TaskDefinitionArn, *subnetId)

	case "ecs-exec":
		cluster := ecs.ListClusters(cfg)
		taskDef := ecs.ListTaskDefinitions(cfg)
		taskDefDetail := ecs.DescribeTaskDefinition(cfg, taskDef)
		taskArn := ecs.ListTasks(cfg, cluster, *taskDefDetail.Family)
		containerName, runtimeId := ecs.DescribeTasks(cfg, cluster, taskArn)
		ecs.ExecuteCommand(cfg, cluster, taskArn, containerName, runtimeId)

	case "StopECSTask":
		cluster := ecs.ListClusters(cfg)
		taskDef := ecs.ListTaskDefinitions(cfg)
		taskDefDetail := ecs.DescribeTaskDefinition(cfg, taskDef)
		taskArn := ecs.ListTasks(cfg, cluster, *taskDefDetail.Family)
		ecs.StopTask(cfg, cluster, taskArn)
	}
}
