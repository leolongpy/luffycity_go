package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 1000
	leftMost := make(chan int)
	right := leftMost
	left := leftMost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftMost)
}
