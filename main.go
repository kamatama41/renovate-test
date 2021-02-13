package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/protobuf/proto"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	cli := s3.New(sess)
	out, err := cli.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	for _, b := range out.Buckets {
		println(*b.Name)
	}
	println(*proto.Bool(true))
}
