package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	argRegion = flag.String("region", "ap-northeast-1", "specify region name")
	argProfile = flag.String("profile", "", "specify credential profile name")
	argBucket = flag.String("bucket", "", "specify bucket name")
	argPrefix = flag.String("prefix", "", "specify object key prefix")
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			flag.Usage()
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	config := aws.Config{Region: argRegion, MaxRetries: aws.Int(10)}
	if *argProfile != "" {
		creds := credentials.NewSharedCredentials("", *argProfile)
		config.Credentials = creds
	}
	sess := session.New(&config)
	svc := s3.New(sess)

	params := &s3.ListObjectsV2Input{Bucket: argBucket, Prefix: argPrefix}
	jst, _ := time.LoadLocation("Asia/Tokyo")
	svc.ListObjectsV2Pages(params,
		func(page *s3.ListObjectsV2Output, lastPage bool) bool {
			for _, obj := range page.Contents {
				fmt.Printf("%s %10d %s\n", obj.LastModified.In(jst).Format("2006-01-02"), *obj.Size, *obj.Key)
			}
			return *page.IsTruncated
		})
}

