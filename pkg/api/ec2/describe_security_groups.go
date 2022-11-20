package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type DescribeSecurityGroupsAPI interface {
	DescribeSecurityGroups(ctx context.Context,
		params *ec2.DescribeSecurityGroupsInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
}

type SecurityGroups struct {
	List []string
}

func (m *SecurityGroups) Set(value string) {
	m.List = append(m.List, value)
}

func describeSecurityGroups(c context.Context, api DescribeSecurityGroupsAPI, input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	return api.DescribeSecurityGroups(c, input)
}

/*
Returns data for the security group for which the specified authentication information is held.

For aws cli -> aws ec2 describe-security-groups
*/
func DescribeSecurityGroups(cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeSecurityGroupsInput{}

	resp, err := describeSecurityGroups(context.TODO(), client, input)
	if err != nil {
		return errFetchSG
	}

	ss := new(SecurityGroups)
	for _, sg := range resp.SecurityGroups {
		fmt.Println(*sg.GroupId)
		fmt.Println(*sg.GroupName)

		ss.Set(*sg.GroupId)
	}

	return nil
}
