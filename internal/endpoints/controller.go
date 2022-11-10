package endpoints

import (
	"awsh/pkg/service/ecs"
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
	case "StartECS":
		err := ecs.StartEcs(cfg)
		return err

	case "ecs-exec":
		err := ecs.EcsExec(cfg)
		return err

	case "StopECSTask":
		err := ecs.StopEcsTask(cfg)
		return err

	default:
		return nil
	}
}
