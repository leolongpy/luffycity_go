package main

import (
	"fmt"
	"os"
)

var (
	AllStudent []*Student
)

// 循环打印帮助信息
func showMenu() {
	fmt.Println("1.添加学生")
	fmt.Println("2.修改学生")
	fmt.Println("3.打印学生")
	fmt.Println("4.结束程序")
}

// 用户输入的方法
func inputStudent() *Student {
	// 定义输入的接收变量
	var (
		username string
		sex      int
		score    float32
		grade    string
	)
	fmt.Println("请输入学生姓名")
	_, _ = fmt.Scanf("%s\n", &username)
	fmt.Println("请输入学生性别:[0|1]")
	_, _ = fmt.Scanf("%d\n", &sex)
	fmt.Println("请输入学生分数:[0-100]")
	_, _ = fmt.Scanf("%f\n", &score)
	fmt.Println("请输入学生年级:")
	_, _ = fmt.Scanf("%s\n", &grade)
	// 创建对象
	stu := NewStudent(username, sex, score, grade)
	return stu
}

// 添加学生的方法
func AddStudent() {
	stu := inputStudent()
	for index, v := range AllStudent {
		if v.Useranme == stu.Useranme {
			AllStudent[index] = stu
			fmt.Println("更新成功")
			return
		}
	}
	AllStudent = append(AllStudent, stu)
	fmt.Println("学生添加成功")
}

// 修改
func ModifyStudent() {
	stu := inputStudent()
	for index, v := range AllStudent {
		if v.Useranme == stu.Useranme {
			AllStudent[index] = stu
			fmt.Println("更新成功")
			return
		}
	}
	AllStudent = append(AllStudent, stu)
	fmt.Println("学生添加成功")
}

func ShowAllStudent() {
	for _, v := range AllStudent {
		fmt.Printf("学生:%s 信息:%#v\n", v.Useranme, v)
	}
	fmt.Println()
}
func main() {
	for {
		showMenu()
		var i int
		_, _ = fmt.Scanf("%d\n", &i)
		switch i {
		case 1:
			AddStudent()
		case 2:
			ModifyStudent()
		case 3:
			ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}
