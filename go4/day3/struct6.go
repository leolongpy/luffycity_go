package main

import "fmt"

//匿名字段
type student4 struct {
	name string
	string
	int
}

func main() {
	var stu1 = student4{
		name: "豪杰",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.string)
	fmt.Println(stu1.int)
}
