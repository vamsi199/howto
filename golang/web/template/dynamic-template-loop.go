package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

type person struct {
	Fname string
	Mname string
	Lname string
	Age   int
}

func main() {
	myr := mux.NewRouter()
	myr.HandleFunc("/", tmplHandler)
	http.ListenAndServe(":8080", myr)
}
func tmplHandler(w http.ResponseWriter, r *http.Request) {
	p1 := person{"Praveen", "Kumar", "K", 37}
	p2 := person{"Srinivasulu", "Reddy", "M", 37}
	peoples := []person{p1, p2}
	for _,p := range peoples {
		tmpl := template.Must(template.ParseFiles("dtpl.gohtml"))
		tmpl.Execute(w, p)
	}
}
