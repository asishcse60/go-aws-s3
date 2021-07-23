package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func DownloadItems(sess *session.Session){
	file, err := os.Create("download1.txt")
	if err != nil{
		log.Fatal(err.Error())
	}
	defer file.Close()
	downloader:=s3manager.NewDownloader(sess)
	_,err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String("go-aws-s3-first-bucket"),
		Key: aws.String("key.txt"),
	})
	if err != nil{
		log.Fatal(err.Error())
	}
	log.Println("successfully downloaded")
}


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


func ListItems(sess *session.Session){
	svc:=s3.New(sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("go-aws-s3-first-bucket"),
	})
	if err != nil{
		log.Fatal(err.Error())
	}
	for _,item:= range resp.Contents{
		fmt.Println("Name:      ", *item.Key)
		fmt.Println("Last modified:      ", *item.LastModified)
		fmt.Println("Size:      ", *item.Size)
		fmt.Println("Storage Class:      ", *item.StorageClass)
		fmt.Println(" ")
	}
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
	fmt.Println("upload items")
	ListItems(sess)
	fmt.Println("List items info")
	DownloadItems(sess)
	fmt.Println("download items")
	fmt.Println("delete items")
}
