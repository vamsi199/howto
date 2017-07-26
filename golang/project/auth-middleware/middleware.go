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

var middleware = alice.New(logger, auth)

func logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("### logger begin")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Since(t1)
		fmt.Println("### logger request duration", t2)
		fmt.Println("### logger end")
	})

}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("### auth end")
		//TODO: allow to continue only if oauth authenticated by verifying the token
		next.ServeHTTP(w, r)

		fmt.Println("### auth end")
	})
}
