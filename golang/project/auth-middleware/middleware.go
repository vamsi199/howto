package main

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
	"time"
)

var middleware = alice.New(logger, ValidateToken)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("### logger begin %v\n", r.URL)
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Since(t1)
		fmt.Println("### logger request duration", t2)
		fmt.Println("### logger end")
	})

}
