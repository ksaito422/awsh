package ecs

import (
	"awsh/pkg/prompt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ecsClustersName struct {
	List []string
}

func (m *ecsClustersName) Set(value string) {
	m.List = append(m.List, value)
}

// Receives a value of type ListClustersOutput in the argument and returns cluster arn in string.
func SelectClusterArn(input *ecs.ListClustersOutput) string {
	// TODO: 引数のnullチェック入れる
	ls := new(ecsClustersName)
	for _, cluster := range input.ClusterArns {
		ls.Set(cluster)
	}

	cluster := prompt.ChooseValueFromPromptItems("Select ECS Clusters", ls.List)
	return cluster
}
