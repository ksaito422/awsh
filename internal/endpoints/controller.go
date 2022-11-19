package endpoints

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/xerrors"
)

// Routing of operation actions on AWS resources.
func (r *Route) Controller(cfg aws.Config, action Operation) error {
	switch action {
	// S3
	case ListBuckets:
		if err := r.S3.ListBuckets(cfg); err != nil {
			return xerrors.Errorf("%w", err)
		}
		return nil

	case ListObjects:
		if err := r.S3.ListObjects(cfg); err != nil {
			return xerrors.Errorf("%w", err)
		}
		return nil

	case DownloadObject:
		if err := r.S3.DownloadObject(cfg); err != nil {
			return xerrors.Errorf("%w", err)
		}
		return nil

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
