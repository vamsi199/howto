//TODO:
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

var mySigningKey = []byte("secret")

var middleware = alice.New(ValidateToken)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: 1: verify the login

	//TODO: 2: if the login failed, exit

	//3: create and return token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	//claims["name"] = "??"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	//TODO return token in body or header?
	w.Write([]byte(tokenString))
}

//func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
func ValidateToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, token.Raw+"\nUnauthorized access to this resource\n"+err.Error())
			http.Redirect(w, r, "localhost:8080/login", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
			http.Redirect(w, r, "localhost:8080/login", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.Handle("/hello", middleware.ThenFunc(welcomeHandler))
	http.HandleFunc("/login", LoginHandler)

	http.ListenAndServe(":8080", nil)
}
