package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("defer")
			runtime.Goexit()
			defer fmt.Println("c.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
}
