/*send the client request with data in body. give examples of sending data in differnt formats listed below:
1) form data (types: text, file)
2) x-www-form-urlencoded
3) raw data: json or any plain text formats
4) binary
5) multipart file attachment????? */

package main

import (
	"net/http"

	"bytes"
	"mime/multipart"
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

func main(){
	var b bytes.Buffer
	content:= "hello"
	b.WriteString(content)
	w:=multipart.NewWriter(&b)
	file,err:=os.Open("hello.txt")
	if err!=nil {
		fmt.Println("cannot open the file",err)
	}
	defer file.Close()
	file.WriteString(content)
	iow,err:=w.CreateFormFile("value","hello.txt")
	if err!= nil{
		fmt.Println("cannot create form file",err)
	}
	_,err=io.Copy(iow,file)
	if err!= nil{
		fmt.Println("cannot copy to form file",err)
	}
	keywriter,err:=w.CreateFormField("key")
	if err!= nil{
		fmt.Println("cannot create key",err)
	}
	_,err=	keywriter.Write([]byte("hello"))
	if err!= nil{
		fmt.Println("cannot create key",err)
	}
	err=w.Close()
	if err!= nil{
		fmt.Println("cannot close multipart writer",err)
	}


	url := "https://localhost:8080/formdatafile"
	req,err := http.NewRequest("POST",url,bytes.NewBufferString(content))
	if err!= nil{
		fmt.Println("cannot set the request",err)
	}
	req.Form.Set("hello","hello.txt")
	client:=&http.Client{}
	resopnse,err:=client.Do(req)
	if err!= nil{
		fmt.Println("cannot get the responce",err)
	}
	defer resopnse.Body.Close()
	byte,err:=ioutil.ReadAll(resopnse.Body)
	if err!= nil{
		fmt.Println("cannot read the responce",err)
	}
	fmt.Println(string(byte))



}
