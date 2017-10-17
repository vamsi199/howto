package fb

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"os"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& handleFacebookLogin begin")
	fmt.Println("app url:", r.Host)

	redirect_uri := "http://" + r.Host + "/callback"

	values := url.Values{}
	values.Add("client_id", os.Getenv("FB_ID"))
	values.Add("redirect_uri", redirect_uri)
	values.Add("response_type", "code")
	values.Add("scope", "public_profile")

	id := uuid.New()
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://www.facebook.com/dialog/oauth?%s", values.Encode())
	fmt.Println("redirectRequestUrl", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}

//https://graph.facebook.com/me?
//https://www.facebook.com/dialog/oauth
