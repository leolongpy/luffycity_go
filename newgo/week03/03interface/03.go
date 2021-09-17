package main

import "fmt"

func main() {
	var s interface{} = false
	s1, ok := s.(string)
	fmt.Println(s1, ok)

	s2, ok := s.(int)
	fmt.Println(s2, ok)
	switch s.(type) {
	case string:
		fmt.Println("是个string")
	case int:
		fmt.Println("是个int")
	default:
		fmt.Println("未知的type")
	}
}
