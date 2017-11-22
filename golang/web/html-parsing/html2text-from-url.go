// based on example from https://github.com/jaytaylor/html2text

package main

import (
	"fmt"
	"github.com/jaytaylor/html2text"
	"net/http"
)

func main() {
	url := "https://www.blog.google/products/maps/google-maps-gets-new-look/"
	text, err := html2textFromUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}

func html2textFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	text, err := html2text.FromReader(resp.Body, html2text.Options{PrettyTables: false})
	if err != nil {
		return "", err
	}
	return text, nil
}
