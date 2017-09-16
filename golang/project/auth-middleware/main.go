package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/hello", middleware.ThenFunc(welcomeHandler))

	http.ListenAndServe(":8080", nil)

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("### welcomeHandler begin")
	w.Write([]byte("hello world"))
	fmt.Println("### welcomeHandler end")
}
