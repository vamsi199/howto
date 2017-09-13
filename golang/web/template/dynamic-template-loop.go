
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

var persons []person
var tpl string

func main() {
	p1 := person{"Praveen", "Kumar", "K", 36}
	p2 := person{"Srinivasulu", "Reddy", "M", 36}
	persons = []person{p1, p2}
	myr := mux.NewRouter()
	myr.HandleFunc("/", tmplHandler)
	http.ListenAndServe(":8080", myr)
}
func tmplHandler(w http.ResponseWriter, r *http.Request) {
	for x,_:=range persons {
		tpl = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>dynamic-template1</title>
		</head>
		<body>
			Name :` + persons[x].Fname + ` ` + persons[x].Mname + `.` + persons[x].Lname + `<br/>
			Age :` + strconv.Itoa(persons[x].Age) + `<br/>
		</body>
		</html>
		`
		fmt.Fprintf(w, tpl)
	}
}
