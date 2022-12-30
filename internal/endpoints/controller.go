package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/xerrors"
)

// Routing of operation actions on AWS resources.
func (r *Route) Controller(cfg aws.Config, action Operation) error {
	switch action {
	// S3
	case ListBuckets:
		if err := r.S3.ListBuckets(cfg); err != nil {
			return err
		}
		return nil

	case ListObjects:
		if err := r.S3.ListObjects(cfg); err != nil {
			return err
		}
		return nil

	case DownloadObject:
		if err := r.S3.DownloadObject(cfg); err != nil {
			return err
		}
		return nil

	// ECS
	case ECS_EXEC:
		if err := r.ECS.EcsExec(cfg); err != nil {
			return err

		}
		return nil

	default:
		return xerrors.New("予期せぬ条件に一致しました")
	}
}
