package main

import (
	"fmt"
	"os"
)

//学生结构体
type Student struct {
	name  string
	age   int
	id    int64
	class string
}

// NewStudent 是一个创造新学生对象构造函数
func NewStudent(name string, age int, id int64, class string) *Student {
	return &Student{
		name:  name,
		age:   age,
		id:    id,
		class: class,
	}
}

// StudetnMgr 定义一个学生信息管理的结构体
type StudentMgr struct {
	AllStudents []*Student
}

// NewStudentMgr 创建新的学生信息管理结构体对象
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		AllStudents: make([]*Student, 0, 100),
	}
}

// AddStudent 添加学生的方法
func (s *StudentMgr) AddStudent(stu *Student) {
	for _, v := range s.AllStudents {
		if v.name == stu.name {
			fmt.Printf("姓名为%s的学生已存在\n", stu.name)
			return
		}
	}
	s.AllStudents = append(s.AllStudents, stu)
}

// EditStudent 修改学生的方法
func (s *StudentMgr) EditStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			s.AllStudents[index] = stu
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}

// DeleteStudent 删除学生的方法
func (s *StudentMgr) DeleteStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			s.AllStudents = append(s.AllStudents[:index], s.AllStudents[index+1:]...)
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}

// ShowStudent 展示学生的方法
func (s *StudentMgr) ShowStudent() {
	for _, v := range s.AllStudents {
		fmt.Printf("name : %s age :%d id :%d calss :%s\n", v.name, v.age, v.id, v.class)
	}
}

func showMenu() {
	fmt.Println("学生信息管理系统！")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 修改学生")
	fmt.Println("3. 删除学生")
	fmt.Println("4. 展示学生")
	fmt.Println("5. 退出")
}
func userInput() *Student {
	var (
		name  string
		age   int
		id    int64
		class string
	)
	fmt.Println("请按提示录入信息")
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("请输入学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入班级：")
	fmt.Scanln(&class)
	newStu := NewStudent(name, age, id, class)
	return newStu
}
func main() {
	stuMgr := NewStudentMgr()
	for {
		showMenu()
		// 获取用户输入的指令
		var i int
		fmt.Print("请输入指令：")
		fmt.Scanln(&i)
		fmt.Println("输入的选项是：", i)
		switch i {
		case 1:
			// 添加学生
			newStu := userInput()
			stuMgr.AddStudent(newStu)
		case 2:
			// 修改学生
			newStu := userInput()
			stuMgr.EditStudent(newStu)
		case 3:
			// 删除学生
			newStu := userInput()
			stuMgr.DeleteStudent(newStu)
		case 4:
			// 展示学生
			stuMgr.ShowStudent()
		case 5:
			// 退出
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}
}
