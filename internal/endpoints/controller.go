package endpoints

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Routing of operation actions on AWS resources.
func (r *Route) Controller(cfg aws.Config, action Operation) error {
	switch action {
	// S3
	case ListBuckets:
		err := r.S3.ListBuckets(cfg)
		return err

	case ListObjects:
		err := r.S3.ListObjects(cfg)
		return err

	case DownloadObject:
		err := r.S3.DownloadObject(cfg)
		return err

	// ECS
	case StartECS:
		err := r.ECS.StartEcs(cfg)
		return err

	case ECS_EXEC:
		err := r.ECS.EcsExec(cfg)
		return err

	case StopECSTask:
		err := r.ECS.StopEcsTask(cfg)
		return err

	default:
		return errors.New("予期せぬ条件に一致しました")
	}
}
