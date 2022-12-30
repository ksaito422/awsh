package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSListClustersAPI interface {
	ListClusters(ctx context.Context,
		params *ecs.ListClustersInput,
		optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
}

func GetAllClusters(c context.Context, api ECSListClustersAPI, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	return api.ListClusters(c, input)
}

// Returns data for the selected ecs cluster.
// For aws cli -> aws ecs list-clusters
func (s *ecsApi) ListClusters(cfg aws.Config) (*ecs.ListClustersOutput, error) {
	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListClustersInput{}

	resp, err := GetAllClusters(context.TODO(), client, input)
	if err != nil {
		return nil, errGetAllCluster
	}

	return resp, nil
}
