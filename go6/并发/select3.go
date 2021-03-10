package main

import (
	"fmt"
	"time"
)

func main() {
	//创建管道
	output1 := make(chan string ,10)
	//子协程写数据
	go write1(output1)
	//取数据
	for s :=range output1{
		fmt.Println("res:",s)
		time.Sleep(time.Second)
	}
}
func write1(ch chan string)  {
	for  {
		select {
		case ch<-"hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond*500)
	}
}
