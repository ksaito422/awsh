package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	
	"awsh/pkg/prompt"
)

type ECSListClustersAPI interface {
	ListClusters(ctx context.Context,
		params *ecs.ListClustersInput,
		optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
}

type ECSClustersName struct {
	List []string
}

func (m *ECSClustersName) Set(value string) {
	m.List = append(m.List, value)
}

func GetAllClusters(c context.Context, api ECSListClustersAPI, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	return api.ListClusters(c, input)
}

// aws ecs list-clusters
func ListClusters(cfg aws.Config) string {
	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListClustersInput{}

	result, err := GetAllClusters(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving clusters:")
		fmt.Println(err)
		return "error"
	}

	ss := new(ECSClustersName)
	for _, cluster := range result.ClusterArns {
		ss.Set(cluster)
	}

	select_cluster := prompt.ChooseValueFromPromptItems("Select ECS Clusters", ss.List)

	return select_cluster
}
