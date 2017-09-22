// dynamic template that uses a struct to inject the data and render it.
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

const dtmpl = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>dynamic-template1</title>
</head>
<body>
    Name : {{.Fname}} {{.Mname}}.{{.Lname}}<br/>
    Age : {{.Age}}<br/>
</body>
</html>
	`

func main() {
	p := person{"Praveen", "Kumar", "K", 36}
	tmpl, err := template.New("dynamic").Parse(dtmpl)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, p)
	})
	http.ListenAndServe(":8080", nil)
}
