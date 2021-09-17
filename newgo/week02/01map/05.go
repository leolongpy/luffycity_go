package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"log"
	"time"
)

func main() {
	m := cmap.New()
	//循环写
	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()
	// 循环读

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			v, exists := m.Get(key)
			if exists {
				log.Printf("[%s=%v]", key, v)
			}
		}
	}()

	time.Sleep(1 * time.Hour)
}
