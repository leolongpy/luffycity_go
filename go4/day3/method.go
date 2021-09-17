package main

import "fmt"

type people struct {
	name   string
	gender string
}

func (p *people) dream() {
	p.gender = "男"
	fmt.Printf("%s的梦想是不用上班也有钱拿！\n", p.name)
}
func main() {
	var haojie = &people{
		name:   "豪杰",
		gender: "爷们",
	}
	haojie.dream()
	fmt.Println(haojie.gender)
}
