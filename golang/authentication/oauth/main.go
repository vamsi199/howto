package main

import (
	"github.com/muly/howto/golang/authentication/oauth/fb"
	"github.com/muly/howto/golang/authentication/oauth/github"
	"github.com/muly/howto/golang/authentication/oauth/gitlab"
	"github.com/muly/howto/golang/authentication/oauth/google"
	"github.com/muly/howto/golang/authentication/oauth/linkedin"
	"io"
	"net/http"
)

const loginhtml = `<!DOCTYPE html>
<html>
<head></head>
<body>
<p><a href="/facebooklogin">LOGIN WITH FACEBOOK </a></p>
<p><a href="/githublogin">LOGIN WITH GITHUB</a></p>
<p><a href="/gitlablogin">LOGIN WITH GITLAB</a></p>
<p><a href="/googlelogin">LOGIN WITH GOOGLE</a></p>
<a href="/linkedinlogin">LOGIN WITH LINKEDIN</a>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/githublogin", github.HandleLogin)
	http.HandleFunc("/gitlablogin", gitlab.HandleLogin)
	http.HandleFunc("/facebooklogin", fb.HandleLogin)
	http.HandleFunc("/linkedinlogin", linkedin.HandleLogin)
	http.HandleFunc("/googlelogin", google.HandleLogin)
	http.HandleFunc("/googlecallback", google.HandleCallback)
	http.HandleFunc("/googlecallbacktoken", google.HandleCallbackToken)

	http.ListenAndServe(":8081", nil)
}

func handleindex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}
