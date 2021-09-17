package main

import (
	"fmt"
	"sync"
	"time"
)

//解决map线程不安全，自己加锁

type concurrentMap struct {
	mp map[int]int
	sync.RWMutex
}

func (c *concurrentMap) Set(key, value int) {
	//加写锁
	c.Lock()
	c.mp[key] = value
	c.Unlock()
}
func (c *concurrentMap) Get(key int) int {
	//获取读锁
	c.RLock()
	res := c.mp[key]
	c.RUnlock()
	return res
}
func main() {
	c := concurrentMap{
		mp: make(map[int]int),
	}
	go func() {
		for i := 0; i < 10000; i++ {
			c.Set(i, i)
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			res := c.Get(i)
			fmt.Printf("[cmap.get][%d=%d]\n", i, res)
		}
	}()
	time.Sleep(1 * time.Hour)
}
