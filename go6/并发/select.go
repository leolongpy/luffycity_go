package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test2"
}
func main() {
	//select {
	//case <-chan1:
	////如果产成功读到数据，则进行该case语句
	//case chan2 <-:
	//	// 如果成功向chan2写入数据，则进行该case处理语句
	//default:
	//	// 如果上面都没有成功，则进入default处理流程
	//}
	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
