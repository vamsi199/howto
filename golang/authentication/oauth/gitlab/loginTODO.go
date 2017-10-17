package gitlab

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& handleGithubLogin begin")
	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)
	//ctx := r.Context()
	//TODO: get session from context

	//redirect_uri := "http://localhost:8080/callback"
	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri:", redirect_uri)

	values := url.Values{}
	values.Add("client_id", "385fa828028158f53e8eb52bfa3112830d4a86d3771a66de2d1d747f9960b701")
	values.Add("redirect_uri", redirect_uri)
	values.Add("response_type", "code")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://gitlab.com/oauth/authorize?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)

	//TODO: save session back to context after saving the id.String()  to a field in session.State

	http.Redirect(w, r, redirectRequestUrl, 302)
}
