package main

import "fmt"

var ch1 chan int       //正常的 可以读 可以写
var ch2 chan<- float64 //只写float64的管道
var ch3 <-chan int     //只读int的管道

func main() {
	//定义通道
	c := make(chan int, 3)
	//转为只写的
	var send chan<- int = c
	//转只读
	var recv <-chan int = c
	send <- 1
	fmt.Println(<-recv)

}
