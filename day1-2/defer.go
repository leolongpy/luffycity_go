package main

import "fmt"

func teseDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("函数马上要结束了")
}

func main() {
	teseDefer()
}
