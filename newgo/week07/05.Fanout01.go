package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch <-chan int, wg *sync.WaitGroup) {
	for {
		task, ok := <-ch
		if !ok {
			break
		}
		time.Sleep(20 * time.Millisecond)
		fmt.Println("启动task", task)
	}
	defer wg.Done()
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	ch := make(chan int)
	for i := 0; i < workers; i++ {
		time.Sleep(1 * time.Millisecond)
		go worker(ch, wg)
	}

	for i := 0; i < tasks; i++ {
		time.Sleep(10 * time.Millisecond)
		ch <- i
	}
	close(ch)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 36)
	wg.Wait()
}
