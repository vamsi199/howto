package main

import (
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
	clientsFile, err := os.OpenFile("myfile", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()


	clients := []*Client{}

	clients = append(clients, &Client{Id: "21", Name: "John", Age: "12"}) // Add clients
	clients = append(clients, &Client{Id: "31", Name: "Fred"})
	clients = append(clients, &Client{Id: "41", Name: "James", Age: "14"})
	clients = append(clients, &Client{Id: "51", Name: "Danny"})
	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}



}
