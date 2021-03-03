package main

import (
	"fmt"
	"time"
)

//测试协程
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i=%d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i=%d\n", i)
		time.Sleep(time.Second)
	}
}
