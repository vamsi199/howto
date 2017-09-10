//TODO: simple static template read from a file, parse and execute it
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

func main() {
	myr := mux.NewRouter()
	myr.HandleFunc("/", tmplHandler)
	http.ListenAndServe(":8080", myr)
}

func tmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("stpl2.gohtml"))
	tmpl.Execute(w,&tmpl)
}