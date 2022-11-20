package ecs

import "golang.org/x/xerrors"

var (
	errDescribeTaskDefinition = xerrors.New("An error occurred while retrieving ECS task definition.")
	errDescribeTask           = xerrors.New("An error occurred while retrieving ECS task.")
	errGetAllCluster          = xerrors.New("An error occurred while retrieving all ECS Cluster.")
	errGetAllTask             = xerrors.New("An error occurred while retrieving all ECS task definiton.")
	errListTask               = xerrors.New("An error occurred while retrieving list ECS task.")
	errRunTask                = xerrors.New("An error occurred while retrieving ECS run task.")
	errStopTask               = xerrors.New("An error occurred while retrieving ECS stop task.")
	errExecuteCommand         = xerrors.New("An error occurred while retrieving execute commang.")
)
