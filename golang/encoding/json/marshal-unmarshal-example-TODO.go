// example program to demonstrate the use of `Marshal`, `Unmarshal` functions of `encoding/json` package

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsondata := `
	{
	"color":"red",
	"breed":"pug",
	"age": 10
	}
	`
	var obj map[string]interface{} //string
	err := json.Unmarshal([]byte(jsondata), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)

	data, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
