//TODO:
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"net/url"
)

var mySigningKey = []byte("secret")

var middleware = alice.New(SetHeader, ValidateToken)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: 1: verify the login

	//TODO: 2: if the login failed, exit

	//3: create and return token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	//claims["name"] = "??"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	//TODO return token in body or header? or in other mode so that the subsequent requests from the same browser will save the token
	//w.Write([]byte(tokenString))
	//w.Header().Add("Authorization","Bearer "+tokenString)

	values := url.Values{}
	values.Add("token", tokenString)

	redirectRequestUrl := fmt.Sprintf("localhost:8080/hello?%s",
		values.Encode())

	http.Redirect(w, r, redirectRequestUrl, 302)
	//w.Write([]byte("authorized"))
}

func SetHeader(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("### SetHeader begin", r.URL)
		//TODO: is this the only way to read the query parameters? or r.FormValue("token")
		vars, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if val, exists := vars["token"]; exists {
			r.Header.Add("Authorization", "Bearer "+val[0])
		}

		next.ServeHTTP(w, r)
		fmt.Println("### SetHeader end")
	})
}

//func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
func ValidateToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("### ValidateToken begin", r.URL)
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})
		if err != nil {
			fmt.Println("### ValidateToken error: ", err)
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprint(w, "\nUnauthorized access to this resource\n"+err.Error())
			http.Redirect(w, r, "localhost:8080/login", 302)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("### ValidateToken error: Token is not valid")
			http.Redirect(w, r, "localhost:8080/login", 302)
			return
		}
		next.ServeHTTP(w, r)

		fmt.Println("### ValidateToken end")
	})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("### welcomeHandler begin")
	w.Write([]byte("hello world"))
	fmt.Println("### welcomeHandler end")
}

func main() {
	http.Handle("/hello", middleware.ThenFunc(welcomeHandler))
	http.HandleFunc("/login", LoginHandler)

	http.ListenAndServe(":8080", nil)
}
