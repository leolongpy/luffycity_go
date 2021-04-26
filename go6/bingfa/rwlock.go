package main

import (
	"fmt"
	"sync"
	"time"
)

//声明读写锁
var rwlock sync.RWMutex
var wg2 sync.WaitGroup

//全局变量
var x2 int

//写数据
func write() {
	rwlock.Lock()
	fmt.Println("write rwlock")
	x2 += 1
	time.Sleep(2 * time.Second)
	fmt.Println("write rwunlock")
	rwlock.Unlock()
	wg2.Done()
}

func read(i int) {
	rwlock.RLock()
	fmt.Println("real rwlock")
	fmt.Printf("gorountine:%d x=%d\n", i, x2)
	time.Sleep(2 * time.Second)
	fmt.Println("read rwunlock")
	rwlock.RUnlock()
	wg2.Done()
}
func main() {
	wg2.Add(1)
	go write()
	time.Sleep(time.Millisecond * 5)
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go read(i)
	}
	wg2.Wait()
}
