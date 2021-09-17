package main

//结构体嵌套
import (
	"log"
)

type Person7 struct {
	Name string
	Age  int
}

type Student7 struct {
	Person7
	StudentId int
}

func (p Person7) SayHello() {
	log.Printf("[Person.SayHello][name:%v]", p.Name)
}

func main() {
	p1 := Person7{
		Name: "leo",
		Age:  123,
	}
	s1 := Student7{
		Person7:   p1,
		StudentId: 99,
	}
	s1.SayHello()
}
