package controller

import (
	"awsh/pkg/ecs"
	"awsh/pkg/s3"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func Main(cfg aws.Config, action string) {
	switch action {
	case "ListBuckets":
		buckets := s3.ListBuckets(cfg)
		for _, v := range buckets {
			fmt.Println(v)
		}

	case "ListObjects":
		buckets := s3.ListBuckets(cfg)
		objects, select_bucket := s3.ListObjects(cfg, buckets)
		fmt.Println("Objects in " + select_bucket + ":")
		for _, item := range objects.Contents {
			fmt.Println("Name:", *item.Key, " | ", "Last modified:", *item.LastModified, " | ", "Size:", item.Size, " | ", "Storage:", item.StorageClass)
		}

	case "GetObject":
		buckets := s3.ListBuckets(cfg)
		objects, select_bucket := s3.ListObjects(cfg, buckets)
		select_object := s3.GetObject(cfg, select_bucket, objects)
		rc := select_object.Body
		defer rc.Close()
		buf := make([]byte, 10000)
		_, err := rc.Read(buf)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s", buf)

	case "DownloadObject":
		buckets := s3.ListBuckets(cfg)
		objects, select_bucket := s3.ListObjects(cfg, buckets)
		s3.DownloadObject(cfg, select_bucket, objects)

	case "StartECS":
		// TODO: タスク定義、クラスター、ネットワーク設定を選択して、関数に渡す
		ecs.StartContainer(cfg)

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
