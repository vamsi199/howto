package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.Handle("/", middleware.ThenFunc(welcomeHandler))

	http.ListenAndServe(":8080", nil)

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("### welcomeHandler begin")
	w.Write([]byte("hello world"))
	fmt.Println("### welcomeHandler end")
}
