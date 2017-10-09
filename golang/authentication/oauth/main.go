package main

import ("net/http"
	"io"
	"github.com/muly/howto/golang/authentication/oauth/fb"
	"github.com/muly/howto/golang/authentication/oauth/github"
	"github.com/muly/howto/golang/authentication/oauth/linkedin"
)


const loginhtml = `<!DOCTYPE html>
<html>
<head></head>
<body>
<p><a href="/facebooklogin">LOGIN WITH FACEBOOK </a></p>
<p><a href="/githublogin">LOGIN WITH GITHUB</a></p>
<a href="/linkedinlogin">LOGIN WITH LINKEDIN</a>
</body>
</html>
`
func main() {
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/githublogin", github.HandleLogin)
	http.HandleFunc("/facebooklogin", fb.HandleLogin)
	http.HandleFunc("/linkedinlogin", linkedin.HandleLogin)
	http.ListenAndServe(":8081", nil)
}

func handleindex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}
