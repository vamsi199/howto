// dynamic template that uses a struct to inject the data and render it.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

type person struct {
	Fname string
	Mname  string
	Lname string
	Age int
}

func main() {
	myr := mux.NewRouter()
	myr.HandleFunc("/", tmplHandler)
	http.ListenAndServe(":8080", myr)
}
func tmplHandler(w http.ResponseWriter, r *http.Request) {
	p := person{"Praveen", "Kumar","K",37}
	tmpl := template.Must(template.ParseFiles("dtpl.gohtml"))
	tmpl.Execute(w, p)
}
