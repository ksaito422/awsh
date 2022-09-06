package ec2

import (
	"awsh/pkg/prompt"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type DescribeSubnetsAPI interface {
	DescribeSubnets(ctx context.Context,
		params *ec2.DescribeSubnetsInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
}

type Subnets struct {
	List []string
}

func (m *Subnets) Set(value string) {
	m.List = append(m.List, value)
}

func describeSubnets(c context.Context, api DescribeSubnetsAPI, input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	return api.DescribeSubnets(c, input)
}

/*
Returns data for the selected subnet.

For aws cli -> aws ec2 describe-subnets
*/
func DescribeSubnets(cfg aws.Config) *string {
	client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeSubnetsInput{}

	resp, err := describeSubnets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving describe subnets:")
		fmt.Println(err)
		os.Exit(1)
	}

	ss := new(Subnets)
	for i, subnet := range resp.Subnets {
		for _, tag := range subnet.Tags {
			if *tag.Key == "Name" {
				ss.Set(strconv.Itoa(i) + ". " + *tag.Value)
			}
		}
	}

	// Name tagからサブネットIDを取得
	tag := prompt.ChooseValueFromPromptItems("Select subnet tags", ss.List)
	cnt := strings.Index(tag, ".")
	i, _ := strconv.Atoi(tag[0:cnt])
	return resp.Subnets[i].SubnetId
}
