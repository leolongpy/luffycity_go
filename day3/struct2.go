package main

import "fmt"

type student struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	var haojie = student{
		name: "豪杰",
		age: 19,
		gender: "男",
		hobby: []string{"篮球","足球","双色球"},
	}

	fmt.Println(haojie)
	fmt.Println(haojie.name)
	fmt.Println(haojie.age)
	fmt.Println(haojie.gender)
	fmt.Println(haojie.hobby)

	var wangzhan = student{}
	fmt.Println(wangzhan.name)
	fmt.Println(wangzhan.age)
	fmt.Println(wangzhan.gender)
	fmt.Println(wangzhan.hobby)

	var yawei = new(student)
	fmt.Println(yawei)
	yawei.name = "亚伟"
	yawei.age = 18
	fmt.Println(yawei.name, yawei.age)

	var nazha = &student{}
	fmt.Println(nazha)
	nazha.name="222"
	fmt.Println(nazha.name)

	var stu1 = student{
		"豪杰",
		18,
		"男",
		[]string{"男人", "女人"},
	}
	fmt.Println(stu1.name, stu1.age)

	//键值对初始化
	var stu2 = &student{
		name:   "豪杰",
		gender: "男",
	}
	fmt.Println(stu2.name, stu2.age, stu2.gender)

}
