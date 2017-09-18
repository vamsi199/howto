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
	// configuration fields //TODO: crete single object to store all the required config fields for s3. i.e the list below
	accessId := os.Getenv("s3accessid")
	accessKey := os.Getenv("secretaccesskey")
	token := ""
	region := "us-east-1"
	bkt = aws.String("jenkins19")

	cred := credentials.NewStaticCredentials(accessId, accessKey, token)
	_, err := cred.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	conf := aws.NewConfig()
	conf.Credentials = cred
	conf.Region = aws.String(region)

	sess = session.Must(session.NewSession(conf))

}

func main() {

	key, err := upload("file")
	if err != nil {
		fmt.Println("upload error:", err)
	}
	fmt.Printf("uploaded %v sucessfully", key)

	keys, err := list()
	if err != nil {
		fmt.Println("list error:", err)
	}
	fmt.Printf("list sucessful %v", keys)

	for _, key := range keys {
		if err = download(key); err != nil {
			fmt.Printf("download %v error: %v", key, err)
			continue
		}
		fmt.Println("download sucessfull:", key)
	}

	for _, key := range keys {
		if err = delete(key); err != nil {
			fmt.Printf("delete %v error: %v", key, err)
			continue
		}
		fmt.Println("delete sucessfull:", key)
	}

}

func upload(filePath string) (key string, err error) {
	upload := s3manager.NewUploader(sess)
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	key = filePath
	uploadInput := s3manager.UploadInput{Bucket: bkt, Body: f, Key: aws.String(key)}
	_, err = upload.Upload(&uploadInput)
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
