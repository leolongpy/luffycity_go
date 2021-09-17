package main

import "fmt"

func main() {
	m := make(map[string]int)
	keys := make([]string, 0)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key_%d", i)
		keys = append(keys, key)
		m[key] = i
	}
	fmt.Println(m)
	fmt.Println("无序遍历")
	// range 遍历
	for k, v := range m {
		fmt.Printf("[%s=%d]\n", k, v)
	}
	fmt.Println("有序遍历")
	for _, k := range keys {
		fmt.Printf("[%s=%d]\n", k, m[k])
	}
}
