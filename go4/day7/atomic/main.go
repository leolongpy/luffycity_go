package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

//普通版加函数
func add() {
	x++
	wg.Done()
}

// 互斥锁加函数
func mutexadd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

//原子操作
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
		//go mutexadd()
		//go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))

}
