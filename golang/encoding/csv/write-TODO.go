// generate a csv file with the data within a slice of struct
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("myfile")
	if err != nil {
		fmt.Println(err)
	}

	f, err = os.OpenFile("myfile", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	_, err=f.WriteString("hello")
	if err != nil {
		fmt.Println("Write error:",err)
	}

}
