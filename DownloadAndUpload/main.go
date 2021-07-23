package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func UploadItem(sess *session.Session){
	f, err := os.Open("additional.txt")
	if err != nil{
		log.Fatal("could not open file")
	}
	defer f.Close()

	uploader:=s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL: aws.String("public-read"),
		Bucket: aws.String("go-aws-s3-first-bucket"),
		Key: aws.String("key.txt"),
		Body: f,
	})
	if err != nil{
		log.Fatal(err.Error())
	}
	log.Printf("uploaed file :%+v\n", result)
}
func main() {
	fmt.Println("Listening aws buckets...")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err !=nil{
		log.Fatal(err.Error())
	}
	UploadItem(sess)
}
