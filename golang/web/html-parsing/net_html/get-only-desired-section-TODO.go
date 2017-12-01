package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

//sequence:
// Type: ElementNode
// Data: h2
// Attr.Key: id
// Attr.Val: pkg-overview

var requiredNodePath *html.Node

func init() {

	//n := html.Node{}
	requiredNodePath.Type = html.ElementNode
	requiredNodePath.Data = "h2"
	a := html.Attribute{}
	a.Key = "id"
	a.Val = "pkg-overview"
	requiredNodePath.Attr = append(requiredNodePath.Attr, a)

	//requiredNodePath = append(requiredNodePath, n)

}

func main() {
	url := `https://godoc.org/golang.org/x/oauth2`

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	FindMatchingNode(n, requiredNodePath, isMatchingNode)

}

func isMatchingNode(n *html.Node, matchWith *html.Node) bool {
	if n.Type == matchWith.Type && n.Data == matchWith.Data {
		matchCnt := 0
		for _, a := range matchWith.Attr {
			for _, w := range n.Attr {
				if a.Key == w.Key && a.Val == w.Val {
					matchCnt++
				}
			}
		}
		if matchCnt != len(matchWith.Attr) {
			return false
		}
		return true
	}
	return false

}

type WalkFunc func(*html.Node, *html.Node) bool

func FindMatchingNode(root *html.Node, match *html.Node, walkF WalkFunc) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if walkF(n, match) {
			fmt.Println(nodeTypeText(n.Type), n.Data, n.Attr)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(root)
}

func nodeTypeText(nodeType html.NodeType) string {
	switch nodeType {
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
