package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
	"net/http"
	"net/url"
	"os"
	//oauth2 "google.golang.org/api/oauth2/v2"
)

var (
	UrlTokenInfo = "https://www.googleapis.com/oauth2/v3/tokeninfo" //?id_token=XYZ123
	UrlAuth      = google.Endpoint.AuthURL
	UrlToken     = google.Endpoint.TokenURL
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
}

type tokenInfo struct {
	Picture    string `json:"picture"`
	Audience   string `json:"aud"` // must match with client_id
	FamilyName string `json:"family_name"`
	Issuer     string `json:"iss"` // must be one of these: accounts.google.com, https://accounts.google.com
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	Expiration string `json:"exp"`
	Subject    string `json:"sub"` // unique user id
	//Algorithm  string `json:"alg"`
	//AtHash     string `json:"at_hash"`
	//Azp        string `json:"azp"`
	//Iat        string `json:"iat"`
	//Locale     string `json:"locale"`
	//Kid        string `json:"kid"`
}

func Handlers() {
	http.HandleFunc("/googlelogin", HandleLogin)
	http.HandleFunc("/googlecallback", HandleCallback)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& google.HandleLogin begin")
	fmt.Println("app url:", r.Host)

	// 1. Create an anti-forgery state token
	id := uuid.New()

	// 2. send auth request to google
	redirect_uri := "http://" + r.Host + "/googlecallback"
	values := url.Values{}
	values.Add("client_id", os.Getenv("GoogleOauth2ClientId"))
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "profile") //"https://www.googleapis.com/auth/admin.directory.device.mobile.readonly")
	values.Add("response_type", "code")
	values.Add("state", id.String())
	redirectRequestUrl := fmt.Sprintf("%v?%s", UrlAuth, values.Encode())
	fmt.Println("redirectRequestUrl", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("&&& google.HandleCallback begin %v\n", r.URL)

	vals := r.URL.Query()
	code, exists := vals["code"]
	if !exists || len(code) == 0 {
		fmt.Println("ERROR: code missing")
		return
	}



	// 3. TODO: validate the scope

	// 4. exchange code for token
	redirectUri := "http://" + r.Host + "/googlecallback"
	t := getAccessToken(code[0], redirectUri)

	// 5. get user info
	u, err := getUserInfo(t)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// 5.1: validate token
	validateToken(u)

	// 6. TODO authenticate: check if the user already exists in the database

}

func getAccessToken(code string, redirect_uri string) string { // todo: return err

	values := url.Values{}
	values.Add("client_id", os.Getenv("GoogleOauth2ClientId"))
	values.Add("client_secret", os.Getenv("GoogleOauth2ClientSecret"))
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "profile")

	resp, err := http.Post(UrlToken, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(values.Encode())))
	if err != nil {
		fmt.Println("ERROR:", err)
		return ""
	}
	defer resp.Body.Close()

	t := tokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		fmt.Println("ERROR:", err)
		return ""
	}
	return t.IdToken

}

func getUserInfo(token string) (tokenInfo, error) {
	url := fmt.Sprintf("%v?id_token=%v", UrlTokenInfo, token)
	resp, err := http.Get(url)
	if err != nil {
		return tokenInfo{}, err
	}
	defer resp.Body.Close()

	t := tokenInfo{}
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		fmt.Println("ERROR:", err)
		return tokenInfo{}, err
	}

	fmt.Println(t)

	return t, nil
}

func validateToken(t tokenInfo) bool{ //TODO

	//Audience   string `json:"aud"` // must match with client_id
	//Issuer     string `json:"iss"` // must be one of these: accounts.google.com, https://accounts.google.com

	return false
}

//// Resources ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// steps. check "Server flow" section: https://developers.google.com/identity/protocols/OpenIDConnect
// oauth playground: https://developers.google.com/oauthplayground/
// https://developers.google.com/identity/sign-in/web/backend-auth
// scope values: https://developers.google.com/identity/protocols/googlescopes
// https://developers.google.com/api-client-library/
// 	https://github.com/google/google-api-go-client
