package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//带过期时间的map
type Cache struct {
	sync.RWMutex
	mp map[string]*item
}
type item struct {
	value int
	ts    int64
}

func (c *Cache) Get(key string) *item {
	c.RLock()
	defer c.RUnlock()
	return c.mp[key]
}

func (c *Cache) CacheNum() int {
	c.RLock()
	keys := make([]string, 0)
	for k, _ := range c.mp {
		keys = append(keys, k)
	}
	c.RUnlock()
	return len(keys)
}
func (c *Cache) Set(key string, value *item) {
	c.Lock()
	defer c.Unlock()
	c.mp[key] = value
}

func (c *Cache) Clean(timeDelta int64) {
	for {
		now := time.Now().Unix()
		//待删除的key的切片
		toDelKeys := make([]string, 0)
		//先加读锁，把所有待删除的拿到
		c.RLock()
		for k, v := range c.mp {
			//时间比较
			if now-v.ts > timeDelta {
				toDelKeys = append(toDelKeys, k)
			}
		}
		c.RUnlock()

		c.Lock()
		for _, k := range toDelKeys {
			log.Printf("[删除过期的数据][key:%s]", k)
			delete(c.mp, k)
		}
		c.Unlock()
		time.Sleep(2 * time.Second)
	}
}
func main() {
	c := Cache{
		mp: make(map[string]*item),
	}
	go c.Clean(30)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		ts := time.Now().Unix()
		im := &item{
			value: i,
			ts:    ts,
		}
		log.Printf("[设置缓存][item][key:%s][v:%v]", key, im)
		c.Set(key, im)

	}
	log.Printf("缓存中的数据量:%d", c.CacheNum())
	time.Sleep(33 * time.Second)
	log.Printf("缓存中的数据量:%d", c.CacheNum())
	//更新缓存
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key_%d", i)
		ts := time.Now().Unix()
		im := &item{
			value: i,
			ts:    ts,
		}
		log.Printf("[更新缓存][item][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}
	log.Printf("缓存中的数据量:%d", c.CacheNum())
	select {}

}
