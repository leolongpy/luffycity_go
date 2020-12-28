package main

import "fmt"

//字符串翻转操作
func main() {
	s1 := "hello"
	bytArr := []byte(s1)
	s2 := ""
	for i := len(bytArr) - 1; i >= 0; i-- {
		s2 = s2 + string(bytArr[i])
	}
	fmt.Println(s2)
	//方法2
	length := len(bytArr)
	for i := 0; i < length/2; i++ {
		bytArr[i], bytArr[length-1-i] = bytArr[length-1-i], bytArr[i]
	}
	fmt.Println(string(bytArr))
}
