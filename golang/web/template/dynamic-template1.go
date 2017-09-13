// dynamic template that uses a struct to inject the data and render it.
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type person struct {
	Fname string
	Mname string
	Lname string
	Age   int
}

var p person

func main() {
	p = person{"Praveen", "Kumar", "K", 36}
	myr := mux.NewRouter()
	myr.HandleFunc("/", tmplHandler)
	http.ListenAndServe(":8080", myr)
}
func tmplHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>dynamic-template1</title>
	</head>
	<body>
		Name :` + p.Fname + ` ` + p.Mname + `.` + p.Lname + `<br/>
		Age :` + strconv.Itoa(p.Age) + `<br/>
	</body>
	</html>
	`
	fmt.Fprintf(w, tpl)
}
