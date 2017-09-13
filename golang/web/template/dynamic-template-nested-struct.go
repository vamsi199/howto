package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"strconv"
	"fmt"
)

type person struct {
	Fname string
	Mname string
	Lname string
	Age   int
	Occp
}
type Occp struct {
	Dept  string
	Comp  string
	Ecode string
	Loca  string
}

var persons []person

func main() {
	o1 := Occp{ "Operations", "ABC", "12345", "India"}
	p1 := person{"Praveen", "Kumar", "K", 37,o1}
	o2 := Occp{ "IT", "DEF", "67890", "California"}
	p2 := person{"Srinivasulu", "Reddy", "M", 37,o2}
	persons=[]person{p1,p2}
	myr := mux.NewRouter()
	myr.HandleFunc("/{name}", tmplHandler)
	http.ListenAndServe(":8080", myr)
}
func tmplHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pname := vars["name"]
	p := person{}
	for _, v := range persons {
		if strings.ToLower(v.Fname) == strings.ToLower(pname) {
			p = v
			break
		}
	}
	tpl:=`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>dynamic-template-nested</title>
</head>
<body>
    <b><u>Name :</u></b> `+p.Fname+p.Mname+`.`+p.Lname+`<br/>
    <b><u>Age :</u></b>`+strconv.Itoa(p.Age)+`<br/>
    <b><u>Occupation Details:</u></b><br/>
    Deptartment-`+p.Dept+`<br/>
    Company-`+p.Comp+`<br/>
    Employee Code-`+p.Ecode+`<br/>
    Location-`+p.Loca+`<br/>
</body>
</html>
	`
	fmt.Fprintf(w,tpl)
}
