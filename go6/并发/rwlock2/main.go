package main

import (
	"fmt"
	"sync"
	"time"
)


var rwlock sync.RWMutex
var wg sync.WaitGroup

var x int

func write(){
	for i := 0; i < 100; i++ {
		rwlock.Lock()
		x+=1
		time.Sleep(10*time.Millisecond)
		rwlock.Unlock()
	}
	wg.Done()
}

func read(i int){
	for i:=0;i<100;i++{
		rwlock.RLock()
		time.Sleep(time.Millisecond)
		rwlock.RUnlock()
	}
	wg.Done()
}

func main(){
	start := time.Now().UnixNano()
	wg.Add(1)
	go write()
	for i:=0;i<100;i++{
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
	end:=time.Now().UnixNano()
	fmt.Println("运行时间：",end-start)
}
