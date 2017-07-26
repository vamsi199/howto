package main

import (
	//"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
)

func init() {
	http.HandleFunc("/login", handleIndex)
	http.HandleFunc("/githublogin", handleGithubLogin)
	http.HandleFunc("/callback", handleOauthCallback)

}

// https://developer.github.com/apps/building-integrations/setting-up-and-registering-oauth-apps/about-authorization-options-for-oauth-apps/

const loginhtml = `<!DOCTYPE html>
<html>
<head>
</head>
<body>
<a href="/githublogin">LOGIN WITH GITHUB</a>
</body>
</html>
`

func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}

func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("### handleGithubLogin begin")

	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)


	//redirect_uri := "http://localhost:8080/callback"
	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri:", redirect_uri)

	values := url.Values{}
	values.Add("client_id", "03712bbff7dae4203b4e")
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "user:email")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)

	// save session to context if required
	//TODO: save session back to context after saving the id.String()  to a field in session.State
	//ctx := r.Context()
	//r.WithContext(ctx)

	http.Redirect(w, r, redirectRequestUrl, 302)
}

func handleOauthCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("### handleOauthCallback begin") //TODO: append the request url details to this log entry
	state := r.FormValue("state")
	code := r.FormValue("code")

	fmt.Printf("### handleOauthCallback: state:%v code:%v", state, code)

	//TODO: generate the JWT and add it to response header


	//TODO: then redirect to landing page

	/*
		////ctx := context.WithValue(r.Context(), "state", state)

		ctx := r.Context()
		//TODO: get session from context

		//TODO: compare the state from session with the state from request. if no match, red flag
	*/


	fmt.Println("### handleOauthCallback end")
}
