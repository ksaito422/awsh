package endpoints

import (
	s3ser "awsh/pkg/service/s3"
)

type Route struct {
	Service s3ser.S3Servicer
}

// constructor関数
func NewAppController(r s3ser.S3Servicer) *Route {
	return &Route{Service: r}
}
