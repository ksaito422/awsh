package main

import (
	"awsh/pkg"
	"awsh/pkg/s3"
)

func main() {
	cfg := pkg.Cfg()

	select_bucket := s3.ListBuckets(cfg)
	s3.ListObjects(cfg, select_bucket)
}
