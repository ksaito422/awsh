package main

import (
	"awsh/pkg/config"
	"awsh/pkg/s3"
)

func main() {
	cfg := config.Cfg()

	select_bucket := s3.ListBuckets(cfg)
	s3.ListObjects(cfg, select_bucket)
}
