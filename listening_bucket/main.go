package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func main(){

	fmt.Println("Listening aws buckets...")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err !=nil{
		log.Fatal(err.Error())
	}

	svc := s3.New(sess)
	result, er:= svc.ListBuckets(nil)
	if er !=nil{
		log.Fatal("Error listening buckets...")
	}
	for _,bucket:=range result.Buckets{
		log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
	}

}

