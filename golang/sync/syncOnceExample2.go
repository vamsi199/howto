package main

import (
	"fmt"
	"sync"
	"time"
)

var data string

var onceApps sync.Once

func InitData() {
	onceApps.Do(func() {
		initData()
	})
}

func initData() {
	fmt.Println("Running initData()")
	data = time.Now().String()
}

func main() {
	time.Sleep(1 * time.Second) // added some delay to prove that the time recorded is NOT the default play.golang.org time
	InitData()
	fmt.Println(data) // 2009-11-10 23:00:01

	time.Sleep(10 * time.Second) // added the delay to differentiate the timestamp from the first initData execution
	InitData()
	fmt.Println(data) // again prints 2009-11-10 23:00:01. this proves that the initData() is executed only once
}
