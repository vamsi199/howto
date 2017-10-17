package google

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
	values.Add("client_id", os.Getenv("GOOGLE_ID"))
	values.Add("redirect_uri", redirect_uri)
	values.Add("scope", "https://www.googleapis.com/auth/admin.directory.device.mobile.readonly")
	values.Add("response_type", "code")
	id := uuid.New()
	values.Add("state", id.String())
	redirectRequestUrl := fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?%s", values.Encode())
	fmt.Println("redirectRequestUrl", redirectRequestUrl)
	http.Redirect(w, r, redirectRequestUrl, 302)
}
