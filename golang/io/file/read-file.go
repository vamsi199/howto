package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := `C:\gows\src\github.com\muly\howto\golang\io\file\read-file.go` // windows absolute path
	//filename :=`/c/gows/src/github.com/muly/howto/golang/io/file/read-file.go` // unix absolute path

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
