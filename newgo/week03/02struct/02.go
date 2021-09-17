package main

import "log"

type Person2 struct {
	Name string
	Age  int
}

func main() {
	p1 := Person2{
		Name: "123",
		Age:  123,
	}
	p2 := &p1 // 等同于 var p2 *Person p2 = &p1
	log.Println("结构体中的字段都是值类型，使用&赋值给另外一个，就是浅拷贝")
	p2.Age = 100 //(*p2).Age = 1000
	p1.Name = "456"
	log.Printf("[p1的内存地址:%p][value:%+v]", &p1, p1)
	log.Printf("[p2的内存地址:%p][value:%+v]", p2, *p2)
}
