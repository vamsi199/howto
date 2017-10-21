package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	clientsFile, err := os.Open("myfile")
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []Client{} // Note: []Client{} or []*Client{} will work

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	fmt.Println(clients)

}
