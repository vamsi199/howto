//TODO: GitHub - justinas/alice: Painless middleware chaining for Go
//
//
//
//

package main

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
	"time"
)

var middleware = alice.New(logger)

func logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging begin")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Since(t1)
		fmt.Println("request duration", t2)
		fmt.Println("Logging end")
	})

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {

	http.Handle("/", middleware.ThenFunc(welcomeHandler))

	http.ListenAndServe(":8080", nil)

}
