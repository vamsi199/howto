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
  "github.com/aws/aws-sdk-go/aws/awserr"

)

func main() {
	aws_access_key_id := "AKIAJFLY5AM7GO7R****"                       //please Enter the credentials
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
	fmt.Printf("response %s",awsutil.StringValue(resp))

  //Read or Download object //TODO add new func
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

//List all objects present input Bucket //TODO need to add route
 listinput := &s3.ListObjectsInput{
      Bucket:  aws.String("crudapi"),
      MaxKeys: aws.Int64(2),
  }

  list, err := svc.ListObjects(listinput)
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

  fmt.Println("List of an objects", list)

//Delete objects present in Input //TODO need to add route
  deleteinput := &s3.DeleteObjectInput{
    Bucket: aws.String("crudapi"),
    Key:    aws.String(path),
}

deletedresult, err := svc.DeleteObject(deleteinput)
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

fmt.Println("object deleted sucessfully", deleteinput, deletedresult)


}
