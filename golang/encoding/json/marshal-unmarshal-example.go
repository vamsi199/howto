// example program to demonstrate the use of `Marshal`, `Unmarshal` functions of `encoding/json` package
//// Marshal: struct to json string
//// Unmarshal: json string to struct

package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Color string `json:"color"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func main() {
	marshalExample()
	unmarshalExample()
}

func marshalExample() {
	d := dog{Color: "brown", Breed: "German Shepherd", Age: 5}
	b, _ := json.Marshal(&d)
	fmt.Println(string(b))
}

func unmarshalExample() {
	jsonStr := `{"color":"brown","breed":"German Shepherd","age":5}`
	d := dog{}
	json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(d)
}
