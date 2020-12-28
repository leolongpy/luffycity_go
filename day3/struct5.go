package main

import "fmt"

type student3 struct {
	name string
	age int
	gender string
	hobby []string
}

func newStudent(n string,age int,g string, h []string) *student3 {
	return &student3{
		name: n,
		age: age,
		gender: g,
		hobby: h,
	}
}
func main() {
	hobbySlice := []string{"篮球","球"}
	haojie := newStudent("豪杰", 18, "男", hobbySlice)
	fmt.Println(haojie.name)
}
