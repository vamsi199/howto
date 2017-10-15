package github

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"os"
)

// https://developer.github.com/apps/building-integrations/setting-up-and-registering-oauth-apps/about-authorization-options-for-oauth-apps/

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& github.HandleLogin begin")

	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)

	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri:", redirect_uri)

	values := url.Values{}
	values.Add("client_id", os.Getenv("GITHUB_ID"))
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "user:email")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}

func handleOauthCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.FormValue("state"))
	/*	state := r.FormValue("state")
		////ctx := context.WithValue(r.Context(), "state", state)

		ctx := r.Context()
		//TODO: get session from context


		//TODO: compare the state from session with the state from request. if no match, red flag
	*/
}
