package main

//结构体单例绑定
import (
	"log"
)

type Person8 struct {
	Name string
	Age  int
}

type Student8 struct {
	Person8
	StudentId int
}

func (p Person8) SayHello() {
	log.Printf("[Person.SayHello][name:%v]", p.Name)
}

func (p *Person8) ChangeAge1() {
	p.Age += 10
	log.Printf("[单例绑定方法][Person.ChangeAge1][p.Age:%v]", p.Age)
}

func (p Person8) ChangeAge2() {
	p.Age += 10
	log.Printf("[非指针绑定绑定方法][Person.ChangeAge2][p.Age:%v]", p.Age)
}

func main() {
	p1 := Person8{
		Name: "leo",
		Age:  123,
	}
	s1 := Student8{
		Person8:   p1,
		StudentId: 99,
	}
	s1.SayHello()

	log.Println(s1.Age)
	s1.ChangeAge1()
	log.Println(s1.Age)

	log.Println(s1.Age)
	s1.ChangeAge2()
	log.Println(s1.Age)
}
