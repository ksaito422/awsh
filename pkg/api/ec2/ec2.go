package ec2

import "golang.org/x/xerrors"

var (
	errFetchSG     = xerrors.New("An error occurred while retrieving the security group.")
	errFetchSubnet = xerrors.New("An error occurred while retrieving the subnet.")
)
