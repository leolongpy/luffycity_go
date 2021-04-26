package main

import "fmt"

//生产者 只写
func producter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

//消费者 只读
func cunsumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
func main() {
	c := make(chan int)
	go producter(c)
	cunsumer(c)
	fmt.Println("main结束")
}
