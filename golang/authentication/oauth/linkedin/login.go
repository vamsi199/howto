package linkedin

import (
	"github.com/google/uuid"
	"net/http"
	"fmt"
	"net/url"
	"os"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("&&& linkedin.HandleLogin begin")

	fmt.Println("app url:", r.Host)

	id := uuid.New()
	fmt.Println("id:", id)

	redirect_url := "http://" + r.Host + "/callback"
	fmt.Println("redirect_url:", redirect_url)

	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("client_id", os.Getenv("LINKEDIN_ID"))
	values.Add("redirect_uri", redirect_url)
	values.Add("state", id.String())
	values.Add("scope", "r_basicprofile")

	redirectRequestUrl := fmt.Sprintf("https://www.linkedin.com/uas/oauth2/authorization?%s",
		values.Encode())

	fmt.Println("redirectRequestUrl:", redirectRequestUrl)

	http.Redirect(w, r, redirectRequestUrl, 302)

}
