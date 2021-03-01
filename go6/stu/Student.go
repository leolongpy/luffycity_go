package main

// Student 定义学生结构体
type Student struct {
	Useranme string
	Sex      int
	Score    float32
	Grade    string
}

// 构造方法
func NewStudent(username string, sex int, score float32, grade string) (stu *Student) {
	stu = &Student{
		Useranme: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}
	return
}
func main() {

}
