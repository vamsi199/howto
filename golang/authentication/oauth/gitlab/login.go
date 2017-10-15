package gitlab

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"os"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& gitlab.HandleLogin begin")
	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)

	redirect_uri := "http://" + r.Host + "/callback"
	fmt.Println("redirect_uri:", redirect_uri)

	values := url.Values{}
	values.Add("client_id", os.Getenv("Gitlab_ClientID"))
	values.Add("redirect_uri", redirect_uri)
	values.Add("response_type", "code")
	values.Add("state", id.String())

	redirectRequestUrl := fmt.Sprintf("https://gitlab.com/oauth/authorize?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)

	http.Redirect(w, r, redirectRequestUrl, 302)
}

