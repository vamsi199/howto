// generate a csv file with the data within a slice of struct
// slice of struct to csv file

package main

import (
	/*"os"
	"encoding/csv"
	"strconv"*/
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

type data struct {
	Name  string
	Age   int
	Marks int
}




func main(){
	f, err := os.Create("students")

	if err !=nil{
		return
	}
	f, err = os.OpenFile("students", os.O_WRONLY, os.ModeAppend)
	if err !=nil{
		return
	}
	w := csv.NewWriter(f)
	w.Comma = ','
	//var dat data

	var  x []data
	//x = append(x,data{Name: "name", Age: 0, Marks: 0 })
	x = append(x,data{Name: "sai", Age: 25, Marks: 90 })
	x = append(x,data{Name: "vamsi", Age: 25, Marks: 95 })

	//fmt.Println(x)
	err =w.Write([]string{"name", "age", "marks"})
	if err !=nil{
		fmt.Println(err)
		return
	}
	for _,v := range x{
		//fmt.Println(v)
		//var s[]string
		//s = append(s, v.Name, strconv.Itoa(v.Age), strconv.Itoa(v.Marks))

		err:=w.Write([]string{v.Name, strconv.Itoa(v.Age), strconv.Itoa(v.Marks)})
		if err !=nil{
			fmt.Println(err)
			return
		}
		//fmt.Println(s)
	}
	w.Flush()



}
