package google

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& google.HandleLogin begin")
	fmt.Println("app url:", r.Host)

	redirect_uri := "http://" + r.Host + "/googlecallback"

	values := url.Values{}
	values.Add("client_id",os.Getenv("GOOGLE_ID") )
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "profile")//"https://www.googleapis.com/auth/admin.directory.device.mobile.readonly")
	values.Add("response_type", "code")
	id := uuid.New()
	values.Add("state", id.String())
	redirectRequestUrl := fmt.Sprintf("%v?%s", google.Endpoint.AuthURL, values.Encode())
	fmt.Println("redirectRequestUrl", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}

func HandleCallback(w http.ResponseWriter, r *http.Request){
	fmt.Println("&&& google.HandleCallback begin %v", r.URL)


	vals := r.URL.Query()
	code , exists:= vals["code"]
	if !exists{
		//TODO:
	}

	//TODO: validate the scope

	redirectUri :="http://" + r.Host + "/googlecallbackgettoken"



	http.Redirect(w, r, accessTokenUrl(code[0], redirectUri), 302)

}

func accessTokenUrl(code string, redirect_uri string)string{



	values := url.Values{}
	values.Add("client_id" ,os.Getenv("GOOGLE_ID"))
	values.Add("client_secret",  os.Getenv("Google_Client_secret"))
	values.Add("grant_type" , "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", redirect_uri)

	url := fmt.Sprintf("%v?%s", google.Endpoint.TokenURL, values.Encode())
	fmt.Println("redirectRequestUrl", url)

	return url

}


func HandleCallbackToken(w http.ResponseWriter, r *http.Request){
	fmt.Println("&&& google.HandleCallbackToken begin %v", r.URL)

	b, _:=ioutil.ReadAll(r.Body)

	fmt.Println(string(b))
}


//
// https://developers.google.com/identity/sign-in/web/backend-auth
// scope values: https://developers.google.com/identity/protocols/googlescopes
