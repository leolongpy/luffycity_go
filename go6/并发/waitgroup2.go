package main

import (
	"fmt"
	"sync"
)


func main() {
	//申明等待组
	var wg sync.WaitGroup
	wg.Add(2)
	go func ()  {
		fmt.Println("子协程1")
		wg.Done()
	}()
	go func ()  {
		fmt.Println("子协程2")
		wg.Done()
	}()
	wg.Wait()
}