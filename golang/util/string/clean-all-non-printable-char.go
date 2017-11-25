// input:
// a string
// process:
// replace all non printable characters with a single space characters,
// replace all multiple repeating spaces to single space character
// output:
// return the processed string
// signature:
// func (string)(string)

// cleanup

package main

import (
	"fmt"
)

func main() {
	var text string

	text = `   1
2	3 45`
	fmt.Println(cleanNonPrintChar(text)) // 1 2 3 45

	text = "1 2  3	4\n5\r6        7"
	fmt.Println(cleanNonPrintChar(text)) //1 2 3 4 5 6 7

	text = ""
	fmt.Println(cleanNonPrintChar(text)) //
}

func cleanNonPrintChar(text string) string {

	clean := make([]rune, 0, len(text))
	var prev rune
	for _, t := range text {
		// replace all non printable char with white space
		if t < 32 { // all non printable characters ascii is < 32
			t = 32
		}
		// reduce all repeating white spaces to single white space
		if !(t == 32 && prev == 32) {
			clean = append(clean, t)
			prev = t
		}
	}

	return string(clean)

}
