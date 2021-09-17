package main

import "fmt"

//Cat结构体
type Cat struct{}

func (c Cat) Say() string {
	return "喵喵喵"
}

type Dog struct{}

func (d Dog) Say() string {
	return "汪汪汪"
}

type Pig struct{}

func (p Pig) Say() (r string) {
	r = "哼哼哼"
	return
}

type animal interface {
	Say() string
}

func main() {
	var animalList []animal
	c := Cat{} // 造一个猫
	d := Dog{} // 造一个狗
	p := Pig{} // 造一个猪
	animalList = append(animalList, c, d, p)
	fmt.Println(animalList)
	for _, itme := range animalList {
		ret := itme.Say()
		fmt.Println(ret)
	}

}
