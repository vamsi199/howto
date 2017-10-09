package oauth

import ("net/http"
	"io"
)


const loginhtml = `<!DOCTYPE html>
<html>
<head></head>
<body>
<a href="/facebooklogin">LOGIN WITH FACEBOOK</a>
</body>
</html>
`


func main() {
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/facebooklogin", handleFacebookLogin)
	http.ListenAndServe(":8081", nil)
}

func handleindex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}
