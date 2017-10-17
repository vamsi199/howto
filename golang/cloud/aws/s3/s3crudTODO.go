// connect to blobstore (S3) on AWS and perform read, write, delete, list file operations

package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

var svc *s3.S3
var bucketName string = "crudapi"

func Init() {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc = s3.New(session.New(), cfg) //TODO: session.New() is deprecated
}

func upload(f string) error {

	// upload file
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/downloads/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		return err
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))

	return nil
}

func main() {

	//create a file
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("temp file create failed: %s", err)
	}

	f.Write([]byte("hello"))

	f.Close()

	path := "hello.txt"
	err = upload(path)
	if err != nil {
		fmt.Printf("upload failed: %s", err)
	}

	// Download
	input := &s3.GetObjectInput{
		Bucket: aws.String("crudapi"),
		Key:    aws.String(path),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("\n Sucessfully get the object", result)

	// list
	listInput := &s3.ListObjectsInput{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(2),
	}

	list, err := svc.ListObjects(listInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("List of objects", list)

	// Delete
	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String("crudapi"),
		Key:    aws.String(path),
	}

	deletedResult, err := svc.DeleteObject(deleteInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println("object deleted sucessfully", deleteInput, deletedResult)
}
