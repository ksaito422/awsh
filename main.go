package main

import (
	"awsh/pkg"
)

func main() {
	cfg := pkg.Cfg()

	select_bucket := pkg.S3ListBuckets(cfg)
	pkg.S3ListObjects(cfg, select_bucket)
}
