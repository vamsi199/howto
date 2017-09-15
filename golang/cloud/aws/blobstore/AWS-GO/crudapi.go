package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	aws_access_key_id := "(AKIAJFLY5AM7GO7R****)"                       //please Enter the credentials
	aws_secret_access_key := "apQ9pVjPO7QezyBveFLPiaE2O47+eLnY5e4*****" //please enter the credentials
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open("/users/vamsibottu/Downloads/Vamsi_Bottu EYE.doc")
	if err != nil {
		fmt.Printf("err opening file: %s", err)
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
		Bucket:        aws.String("crudapi"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s", err)
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
}
