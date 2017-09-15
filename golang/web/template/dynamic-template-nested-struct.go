package main

import (
	"net/http"
	"text/template"
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

const dtmpl  = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>dynamic-template-nested</title>
</head>
<body>
    <b><u>Name :</u></b>{{.Fname}}{{.Mname}}.{{.Lname}}<br/>
    <b><u>Age :</u></b>{{.Age}}<br/>
    <b><u>Occupation Details:</u></b><br/>
    Deptartment-{{.Dept}}<br/>
    Company-{{.Comp}}<br/>
    Employee Code-{{.Ecode}}<br/>
    Location-{{.Loca}}<br/>
</body>
</html>
	`

func main() {
	o1 := Occp{ "Operations", "ABC", "12345", "India"}
	p1 := person{"Praveen", "Kumar", "K", 37,o1}

	tmpl,err:=template.New("dyn-tmpl-nest-str").Parse(dtmpl)
	if err!=nil{
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w,p1)
	})
	http.ListenAndServe(":8080", nil)
}
