// example program to demonstrate the use of Encode, Decode functionality of `encoding/json` package
package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type dogs struct {
	Color string
	Breed string
	Age   int
}


func main() {
	dogdata := dogs{Color: "red", Breed: "pug", Age: 10}

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var obj map[string][]string
	err = json.NewDecoder(file).Decode(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	json.NewEncoder(os.Stdout).Encode(dogdata)
	fmt.Println(dogdata)
	//jsondata := `{"color":"red","breed":"pug","age": 10}`
	/*d := map[string]int{"apple": 5, "lettuce": 7}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(d)*/

}
