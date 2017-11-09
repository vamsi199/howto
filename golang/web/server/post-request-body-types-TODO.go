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
)

func main() {
	http.HandleFunc("/formdatafile", handlerFormDataFileType)
	http.ListenAndServe(":8080", nil)
}

//1b) form data file type example
func handlerFormDataFileType(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "file received")
}
