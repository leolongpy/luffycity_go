package main

import "fmt"

type Humaner interface {
	Say()
}

type Peresoner interface {
	Humaner
	sing(ly string)
}

type Student1 struct {
	name  string
	score int
}

func (s Student1) Say() {
	fmt.Printf("Student[%s,%d]\n", s.name, s.score)
}

//func (s )()  {
//
//}

type Teacher struct {
	name  string
	group string
}

func (t Teacher) Say() {
	fmt.Printf("Teacher[%s,%s]诲人不倦\n", t.name, t.group)
}

type MyStr string

func (str MyStr) Say() {
	fmt.Printf("MyStr[%s] 你你你你！\n", str)
}

func WhoSay(i Humaner) {
	i.Say()
}

func main() {
	s := Student1{"szs", 88}
	t := Teacher{"lis", "666"}
	var tmp MyStr = "3333"
	s.Say()
	t.Say()
	tmp.Say()

	WhoSay(s)
	WhoSay(t)
	WhoSay(tmp)
}
