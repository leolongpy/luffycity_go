package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ch <-chan struct{}) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ch:
			break LABEL
		default:

		}
	}
}
func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 5)
	exitChan <- struct{}{}
	wg.Wait()
	fmt.Println("over")

}
