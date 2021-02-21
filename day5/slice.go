package main

import "fmt"

type Person struct {
	name string
	sex string
	age int
}

type Student1 struct {
	Person
	id int
	addr string
}
func main() {
	s1:=Student1{Person{"zs","female",20},1,"bj"}
	fmt.Println(s1)

	s2:=Student1{Person:Person{"zs","female",20}}
	fmt.Println(s2)

	s3:=Student1{Person:Person{name: "zs"}}
	fmt.Println(s3)

}
