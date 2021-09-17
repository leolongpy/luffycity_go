package main

import "fmt"

type student2 struct {
	name string
	age  int8
}

func main() {
	var stu1 = student2{
		name: "豪杰",
		age:  18,
	}
	stu2 := stu1
	stu2.name = "王展"
	fmt.Println(stu1.name)
	fmt.Println(stu2.name)

	stu3 := &stu1
	fmt.Printf("%T\n", stu3)
	(*stu3).name = "哪找"
	fmt.Println(stu1.name, stu2.name, stu3.name)
}
