package main

import (
	"awsh/internal/welcome"
	"awsh/pkg/config"
	"awsh/pkg/s3"
)

func main() {
	welcome.Main()
	cfg := config.Cfg()

	select_bucket := s3.ListBuckets(cfg)
	s3.ListObjects(cfg, select_bucket)
}
