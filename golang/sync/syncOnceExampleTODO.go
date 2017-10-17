package main

//TODO: need to add notes on what this example is demonstrating

import (
	"fmt"
	"sync"
	"time"
)

var onceIdms sync.Once
var i = 0
var j = 0

func getOnce() (int, int) {
	onceIdms.Do(func() {
		i++
		time.Sleep(10000 * time.Millisecond)
	})
	j++

	return i, j
}

func main() {
	test1()
}

func test1() {
	for k := 0; k < 5; k++ {
		go func() {
			p, q := getOnce()
			fmt.Printf("i=%v, j=%v\n", p, q)
		}()
	}

	time.Sleep(11 * time.Second)

}
