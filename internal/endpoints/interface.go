package endpoints

import "github.com/aws/aws-sdk-go-v2/aws"

type AppController interface {
	Controller(aws.Config, Operation) error
	Operation() string
}

type Route struct {
	route AppController
}

// constructor関数
func NewAppController(r AppController) *Route {
	return &Route{route: r}
}
