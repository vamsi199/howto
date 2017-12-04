package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

//const url = `https://www.ziprecruiter.com/jobs/cybercoders-befe3473/senior-software-engineer-golang-0dbe59d6?source=ziprecruiter-jobs-site`
const htmldata = `<!DOCTYPE html1234>
<html567>
<head>
<title>Page Title</title>
</head>
<body>

<h1 a123=123456>This is a Heading1</h1>
<p>This is a paragraph.</p>

<h1 b123=qwerty>This is a Heading2</h1>

</body>
</html567>`

func main() {

	ParseExample()

}

func ParseExample() {
	//s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	htmlReader := strings.NewReader(htmldata)

	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//htmlReader := resp.Body

	doc, err := html.Parse(htmlReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	var f func(*html.Node)
	f = func(n *html.Node) {
		//if n.Type == html.ElementNode && n.Data == "a" {
		//	for _, a := range n.Attr {
		//		if a.Key == "href" {
		//			//fmt.Println("val: ",a.Val)
		//			break
		//		}
		//	}
		//}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fmt.Println(i, c.Data)
			f(c)
			i++
			if i == 2 {
				return
			}
		}
	}
	f(doc)
}

func myParseExample() {

	n, err := html.Parse(strings.NewReader(htmldata))
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(i, c.Type, NodeTypeString(c.Type), c.Data, c.DataAtom)
		i++
	}

}

func NodeTypeString(typ html.NodeType) string {

	switch typ {
	case html.ErrorNode:
		return "ErrorNode"
	case html.TextNode:
		return "TextNode"
	case html.DocumentNode:
		return "DocumentNode"
	case html.ElementNode:
		return "ElementNode"
	case html.CommentNode:
		return "CommentNode"
	case html.DoctypeNode:
		return "DoctypeNode"
	}
	return ""
}

func TokenTypeString(typ html.TokenType) string {

	switch typ {
	case html.ErrorToken:
		return "ErrorToken"
	case html.TextToken:
		return "TextToken"
	case html.StartTagToken:
		return "StartTagToken"
	case html.EndTagToken:
		return "EndTagToken"
	case html.SelfClosingTagToken:
		return "SelfClosingTagToken"
	case html.CommentToken:
		return "CommentToken"
	case html.DoctypeToken:
		return "DoctypeToken"
	}
	return ""
}

func cleanNonPrintChar(text string) string {

	clean := make([]rune, 0, len(text))
	var prev rune
	for _, t := range text {
		// replace all non printable char with white space
		if t < 32 { // all non printable characters ascii is < 32
			t = 32
		}
		// reduce all repeating white spaces to single white space
		if !(t == 32 && prev == 32) {
			clean = append(clean, t)
			prev = t
		}
	}

	return string(clean)

}
