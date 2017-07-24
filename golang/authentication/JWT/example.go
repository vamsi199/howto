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
	claims["name"] = "??"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

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

		if err == nil {
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource")
		}

	})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {

	http.Handle("/", middleware.ThenFunc(welcomeHandler))

	http.ListenAndServe(":8080", nil)

}
