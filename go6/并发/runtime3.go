package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		go fmt.Println(0)
		fmt.Print(1)
	}
}
