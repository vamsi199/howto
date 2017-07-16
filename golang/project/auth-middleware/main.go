package main

import "net/http"


func main(){

	http.Handle("/", middleware.ThenFunc(welcomeHandler))

	http.ListenAndServe(":8080", nil)

}



func welcomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}

