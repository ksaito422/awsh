//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/service/$GOFILE
package ecs

import (
	"awsh/pkg/api/ecs"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ECSServicer interface {
	EcsExec(aws.Config) error
}

type ECSService struct {
	Api ecs.ECSApi
}

// constructor関数
func NewECSService(s ecs.ECSApi) *ECSService {
	return &ECSService{Api: s}
}
