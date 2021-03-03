package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("子协程结束")
		fmt.Println("子协程正在运行")
		c <- 666
	}()

	num := <-c
	fmt.Println("num=", num)
	fmt.Println("main结束")

}
