package main

import "fmt"

import m "luffycity.com/day4/math_pkg"

const Mode = 1

func main() {
	m.Add(100, 200)
	stu := m.Student{Name: "haojie", Age: 18}
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}
