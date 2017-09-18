// connect to blobstore (S3) on AWS and perform read, write, delete, list file operations

package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

var sess *session.Session
var bkt *string

func init() {
	accesId := os.Getenv("s3accessid")       //"AKIAIOHZVBINWEK35PZA"//
	accessKey := os.Getenv("secretacceskey") //"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"//
	token := ""
	cred := credentials.NewStaticCredentials(accesId, accessKey, token)
	conf := aws.NewConfig()
	conf.Credentials = cred
	conf.Region = aws.String("us-east-1")

	sess = session.Must(session.NewSession(conf))

	bkt = aws.String("jenkins19")
}

func main() {

	key, err := upload("file")
	if err != nil {
		fmt.Println("upload error:", err)
	}

	keys, err := list()
	if err != nil {
		fmt.Println("list error:", err)
	}

	for _, key := range keys {
		if err = download(key); err != nil {
			fmt.Printf("download %v error: %v", key, err)
		}
	}

	for _, key := range keys {
		if err = delete(key); err != nil {
			fmt.Printf("delete %v error: %v", key, err)
		}
	}

}

func upload(filePath string) (key string, err error) {
	upload := s3manager.NewUploader(sess)
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	key = filePath
	upoadinput := s3manager.UploadInput{Bucket: bkt, Body: f, Key: aws.String(key)}
	_, err = upload.Upload(&upoadinput)
	if err != nil {
		return "", err
	}
	return key, nil
}

func download(key string) error {

	return errors.New("not implemented")
}

func delete(key string) error {

	return errors.New("not implemented")
}

func list() (keys []string, err error) {

	return nil, errors.New("not implemented")
}
