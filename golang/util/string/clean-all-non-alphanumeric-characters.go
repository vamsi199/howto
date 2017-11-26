// input:
// a string
// process:
// replace all non alphanumeric characters with a single space characters, 
// replace all multiple repeating spaces to single space character
// output:
// return the processed string
// signature:
// func (string)(string,error)

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := ""
	str = "aroot:*:0:0:System Administrator:/root:/bin/sh   replace   multiple   spaces   to   single   space"
	fmt.Printf("modifed String str: %v\n", cleanNonAlphanumChar(str))


	str = "java/j2ee"
	fmt.Printf("modifed String str: %v\n", cleanNonAlphanumChar(str))
}

func cleanNonAlphanumChar(str string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Printf("invalid reg :%v\n", err)
	}
	modStr := reg.ReplaceAllString(str, " ")
	return modStr
}

