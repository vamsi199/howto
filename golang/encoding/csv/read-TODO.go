//TODO: to read from a csv file and save it back to a slice of struct
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"encoding/json"
)

type data struct {
	Name  string `json:"name"`
	Age   int `json:"age"`
	Marks int `json:"marks"`
}

func main() {
	f, err := os.Open("student")
	if err != nil {
		return
	}
	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'
	d, err := r.ReadAll()
	if err != nil {
		return
	}

	var dat data
	var recs []data
	for i, cols := range d {
		if i==0{
			continue
		}
		dat.Name = string(cols[0])
		dat.Age, err = strconv.Atoi(string(cols[1]))
		dat.Marks, err= strconv.Atoi(string(cols[2]))
		recs = append(recs, dat)

	}
	b, err :=json.Marshal(recs)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
}


