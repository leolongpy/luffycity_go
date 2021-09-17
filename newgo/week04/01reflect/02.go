package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}
type Student struct {
	Person
	StudentId  int
	SchoolName string
	IsBaoSong  bool
	Hobbies    []string
	Labels     map[string]string
}

// 非指针方法
func (s Student) GoToSchool() {
	log.Printf("[去上学][sid:%d]", s.StudentId)
}

// 非指针方法
func (s *Student) GoHome() {
	log.Printf("[回家了][sid:%d]", s.StudentId)
}

// 小写方法
func (s Student) baoSong() {
	log.Printf("[竞赛保送][sid:%d]", s.StudentId)
}

func main() {
	s := Student{
		Person:     Person{Name: "leo", Age: 18},
		StudentId:  123,
		SchoolName: "河北工院",
		IsBaoSong:  true,
		Hobbies:    []string{"唱", "跳", "Rap"},
		Labels:     map[string]string{"k1": "v1", "k2": "v2"},
	}
	//获取目标对象
	t := reflect.TypeOf(s)
	log.Printf("[对象类型的名称：%v]", t.Name())
	//值类型
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		//Field代表对象的字段名对象
		key := t.Field(i)
		//通过v.Field(i).Interface获取字段的值
		value := v.Field(i).Interface()
		anonymous := "非匿名"
		if key.Anonymous {
			anonymous = "匿名"
		}
		log.Printf("[%s 字段][第%d字段][字段的名称:%s][字段的类型:%v][字段的值：%v]",
			anonymous,
			i+1,
			key.Name,
			key.Type,
			value,
		)
		// 通过NumMethod获取对象绑定的方法、

	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		log.Printf("[第%d个方法][方法的名称:%s][方法的类型:%v]", i+1, m.Name, m.Type)
	}
}
