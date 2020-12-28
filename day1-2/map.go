package main

import (
	"fmt"
	"sort"
)

func main() {
	//声明map类型
	var m1 map[string]int

	m1 = make(map[string]int, 100)
	m1["leo"] = 100
	m1["long"] = 90

	fmt.Println(m1)

	m2 := map[string]string{
		"haojie": "hehe",
		"yawei":  "heihei",
	}
	fmt.Println(m2)

	// 判断map中是否存在某个键值对
	v, ok := m2["haojie"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k := range m2 {
		fmt.Println(k)
	}
	//删除
	delete(m2, "haojie")
	fmt.Println(m2)

	m3 := map[string]int{
		"haojie":   100,
		"nazha":    80,
		"cuicui":   70,
		"wangxin":  200,
		"qiaojing": 180,
	}
	var keys = make([]string, 0, 10)
	for key := range m3 {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	//排序
	sort.Strings(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, m3[key])
	}
}
