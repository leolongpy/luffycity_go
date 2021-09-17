package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	c := cache.New(30*time.Second, 5*time.Second)
	c.Set("k1", "v1", 31*time.Second)
	res, ok := c.Get("k1")
	fmt.Println(res, ok)
	time.Sleep(time.Second * 32)
	res, ok = c.Get("k1")
	fmt.Println(res, ok)

}
