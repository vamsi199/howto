package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/facebooklogin", handleFacebookLogin)
	http.ListenAndServe(":8081", nil)
}

const loginhtml = `<!DOCTYPE html>
<html>
<head></head>
<body>
<a href="/facebooklogin">LOGIN WITH FACEBOOK</a>
</body>
</html>
`

func handleindex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& handleFacebookLogin begin")
	fmt.Println("app url:", r.Host)
	id := uuid.New()
	fmt.Println("id:", id)
	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri", redirect_uri)
	values := url.Values{}
	values.Add("client_id", "123050871728605")
	values.Add("redirect_uri", redirect_uri)
	values.Add("response_type", "code")
	values.Add("scope", "public_profile")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://www.facebook.com/dialog/oauth?%s", values.Encode())
	fmt.Println("redirectRequestUrl", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}

//https://graph.facebook.com/me?
//https://www.facebook.com/dialog/oauth
