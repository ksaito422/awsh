package ecs

import "github.com/aws/aws-sdk-go-v2/aws"

type ECSServicer interface {
	StartEcs(aws.Config) error
	EcsExec(aws.Config) error
	StopEcsTask(aws.Config) error
}

type ECS struct {
	service ECSServicer
}

// constructor関数
func NewECSService(s ECSServicer) *ECS {
	return &ECS{service: s}
}
