package main

import "fmt"

type speaker interface {
	speak()
}
type mover interface {
	move()
}

//接口嵌套
type animals interface {
	speaker
	mover
}
type cat struct {
	name string
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}
func main() {
	var x animals
	x = cat{name: "花花"}
	x.move()
	x.speak()
}
