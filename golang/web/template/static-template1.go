//TODO: simple static template that just reads the temaplate from a constant parses and executes it
package main

import (
	"os"
	"text/template"
)

func main() {
	stmpl := `<!DOCTYEP=html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>static-template2</title>
<body>
Hello World from Static Template1
</body>
</html>
`
	tmpl, err := template.New("static").Parse(stmpl)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, stmpl)
}
