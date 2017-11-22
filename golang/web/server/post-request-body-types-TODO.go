//accept the request to server with data in body in various ways listed below:
//1) form data (types: text, file)
//2) x-www-form-urlencoded
//3) raw data: json or any plain text formats
//4) binary
//5) multipart file attachment?????
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/formdatafile", handlerFormDataFileType)
	http.ListenAndServe(":8080", nil)
}

//1b) form data file type example
func handlerFormDataFileType(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32<<30)
	dataFile, _, err := r.FormFile("file") // Here file is the key specified in the form while sending the file in form
	if err != nil{
		fmt.Fprintln(w, "r.formfile error:"+err.Error())
		return
	}
	if dataFile == nil{
		fmt.Fprintln(w, "file NOT received")
		return
	}
	defer dataFile.Close()

	processFile(dataFile)

	fmt.Fprintln(w, "file received")
}

func processFile(f io.Reader)error{
	return nil
}
