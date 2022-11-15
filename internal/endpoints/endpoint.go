package endpoints

import (
	"awsh/pkg/service/ecs"
	"awsh/pkg/service/s3"
)

type Route struct {
	S3  s3.S3Servicer
	ECS ecs.ECSServicer
}

// constructor関数
func NewAppController(
	s3 s3.S3Servicer,
	ecs ecs.ECSServicer,
) *Route {
	return &Route{
		S3:  s3,
		ECS: ecs,
	}
}
