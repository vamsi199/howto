// connect to blobstore (S3) on AWS and perform read, write, delete, list file operations

package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

var sess *session.Session
var s3Client *s3.S3

func main() {

	//filePath := "s3crud2TODO.go"
	bkt := "muly123"

	/*	err := createBkt(bkt)
		if err != nil {
			fmt.Println("bucket create error:", err)
		}
		fmt.Printf("bucket %v created sucessfully\n", bkt)*/

	filePath := "s3crud2TODO.go"
	key, err := upload(filePath, bkt)
	if err != nil {
		fmt.Println("upload error:", err)
		return
	}
	fmt.Printf("uploaded %v sucessfully\n", key)

	keys, err := list(bkt)
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	fmt.Printf("list sucessful %v\n", keys)

	for _, key := range keys {
		if err = download(key, bkt); err != nil {
			fmt.Printf("download %v error: %v\n", key, err)
			continue
		}
		fmt.Println("download sucessfull:", key)
	}

	for _, key := range keys {
		if err = delete(key, bkt); err != nil {
			fmt.Printf("delete %v error: %v\n", key, err)
			continue
		}
		fmt.Println("delete sucessfull:", key)
	}

}
