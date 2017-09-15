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
}

const dtmpl  =`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>dynamic-template1</title>
		</head>
		<body>
			Name :{{.Fname}} {{.Mname}}.{{.Lname}}<br/>
			Age :{{.Age}}<br/>
		</body>
		</html>
		`

func main() {
	p1 := person{"Praveen", "Kumar", "K", 36}
	p2 := person{"Srinivasulu", "Reddy", "M", 36}
	persons := []person{p1, p2}

	tmpl, err:=template.New("dynamic-loop").Parse(dtmpl)
	if err!=nil{
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _,v:= range persons{
			p:=v
			tmpl.Execute(w,p)
		}
	})
	http.ListenAndServe(":8080", nil)
}
