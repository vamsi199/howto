// connect to blobstore (S3) on AWS and perform read, write, delete, list file operations

package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var sess *session.Session
var s3Client *s3.S3

func init() {
	// configuration fields //TODO: create single object to store all the required config fields for s3. i.e the list below
	accessId := os.Getenv("s3accessid") //TODO: need to change the credentialsenv variables to AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY
	accessKey := os.Getenv("s3accesskey")
	token := ""
	region := "us-west-1"

	//fmt.Println("access ID, Key: ", accessId, accessKey)

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

/*func createBkt(bkt string)error{

	input := &s3.CreateBucketInput{Bucket: aws.String(bkt)}
	_, err:= s3Client.CreateBucket(input)
	if err != nil{
		return err
	}

	return nil
}*/

func upload(filePath string, bkt string) (key string, err error) {
	upload := s3manager.NewUploader(sess)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error to open file")
		return "", err
	}
	defer f.Close()

	key = filePath
	uploadInput := s3manager.UploadInput{Bucket: aws.String(bkt), Body: f, Key: aws.String(key)}
	_, err = upload.Upload(&uploadInput)
	if err != nil {
		fmt.Println("error uploading file")
		return "", err
	}
	return key, nil
}

func download(key string, bkt string) error {

	return errors.New("not implemented")
}

func delete(key string, bkt string) error {

	return errors.New("not implemented")
}

func list(bkt string) (keys []string, err error) {

	return nil, errors.New("not implemented")
}
