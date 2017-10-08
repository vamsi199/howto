package main

import (
	//"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/twitterlogin", twitterOuth2TokenHandler)
	http.HandleFunc("/oauthcallback", handleOauthCallback)
	http.ListenAndServe(":8080", nil)
}

// https://developer.github.com/apps/building-integrations/setting-up-and-registering-oauth-apps/about-authorization-options-for-oauth-apps/

const loginhtml = `<!DOCTYPE html>
<html>
<head>
</head>
<body>
<a href="/twitterlogin">LOGIN WITH TWITTER</a>
</body>
</html>
`

func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}

func twitterOuth2TokenHandler(w http.ResponseWriter, r *http.Request){
	url:=`https://api.twitter.com/oauth2/token?grant_type=client_credentials`

	req, err := http.NewRequest("POST", url, nil)
	if err != nil{
		fmt.Println("error creating new request:",err)
		http.Error(w, "error creating new request:"+err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "")
	req.SetBasicAuth("K9UrVeqpRHUEz0NnFCQ0LwBwB", "YjbqSKLTRXCNPzXg2QHFRQZ9xb2M5dh5sQOfkUVpr0RSIlnBNI")

	c:= http.Client{}
	res, err:=c.Do(req)
	if err != nil{
		fmt.Println("error with client.Do:",err)
		http.Error(w, "error with client.Do:"+err.Error(), http.StatusInternalServerError)
		return
	}

	data, err:= ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("error reading request_token response:",err)
		http.Error(w, "error reading request_token response:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(data))
}

func twitterRequestTokenHander(w http.ResponseWriter, r *http.Request){
	url:=`https://api.twitter.com/oauth/request_token`

	req, err := http.NewRequest("POST", url, nil)
	if err != nil{
		fmt.Println("error creating new request:",err)
		http.Error(w, "error creating new request:"+err.Error(), http.StatusInternalServerError)
		return
	}

	//authHeader := `OAuth oauth_callback="http%3A%2F%2Flocalhost%2Fcallback%2F",oauth_consumer_key="K9UrVeqpRHUEz0NnFCQ0LwBwB",oauth_signature_method="HMAC-SHA1",oauth_timestamp="1507347040",oauth_nonce="pbMCCk",oauth_version="1.0",oauth_signature="PkrnsOj3zAsofg%2FwQh7tEMDcKOE%3D"`
	authHeader := `OAuth oauth_consumer_key="K9UrVeqpRHUEz0NnFCQ0LwBwB",oauth_signature_method="HMAC-SHA1",oauth_timestamp="1507348523",oauth_nonce="D9TSXm123",oauth_version="1.0",oauth_signature="KbniuhxCrgPF4o1ICwNkAbiASGo%3D"`
	req.Header.Set("Authorization", authHeader)


	c:= http.Client{}
	res, err:=c.Do(req)
	if err != nil{
		fmt.Println("error with client.Do:",err)
		http.Error(w, "error with client.Do:"+err.Error(), http.StatusInternalServerError)
		return
	}

	data, err:= ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("error reading request_token response:",err)
		http.Error(w, "error reading request_token response:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(data))
}

func handleTwitterLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& handleTwitterLogin begin")

	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)
	//ctx := r.Context()
	//TODO: get session from context

	//redirect_uri := "http://localhost:8080/callback"
	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri:", redirect_uri)

	values := url.Values{}
	//values.Add("client_id", "193222610")
	values.Add("consumer_key", "K9UrVeqpRHUEz0NnFCQ0LwBwB")
	values.Add("consumer_secret", "YjbqSKLTRXCNPzXg2QHFRQZ9xb2M5dh5sQOfkUVpr0RSIlnBNI")
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "public_profile")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://api.twitter.com/oauth/authenticate?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)

	//TODO: save session back to context after saving the id.String()  to a field in session.State

	http.Redirect(w, r, redirectRequestUrl, 302)
}

func handleOauthCallback(w http.ResponseWriter, r *http.Request) {
	/*	state := r.FormValue("state")
		////ctx := context.WithValue(r.Context(), "state", state)

		ctx := r.Context()
		//TODO: get session from context

		//TODO: compare the state from session with the state from request. if no match, red flag
	*/
}
