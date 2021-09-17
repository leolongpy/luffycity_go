package main

import "log"

type Person3 struct {
	Name string
	Age  int
}

func main() {
	p1 := new(Person3)

	p1.Name = "leo"
	p1.Age = 123
	p2 := p1
	p1.Age = 19
	p2.Name = "long"
	log.Printf("[p1的内存地址:%p][value:%+v]", p1, *p1)
	log.Printf("[p2的内存地址:%p][value:%+v]", p2, *p2)
}
