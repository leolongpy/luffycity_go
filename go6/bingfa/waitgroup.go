package main

import "fmt"

func main() {
	//创建管道
	ch := make(chan int)
	//计数，代表协程个数
	count := 2
	go func() {
		fmt.Println("子协程1")
		ch <- 1
	}()

	go func() {
		fmt.Println("子协程2")
		ch <- 1
	}()

	for range ch {
		count--
		if count == 0 {
			close(ch)
		}
	}
}
