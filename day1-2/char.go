package main

import "fmt"

func main() {
	//s1 := "Golang"
	//c1 := 'G'
	//fmt.Println(s1,c1)
	s2 := "中国"
	c2 := "中"
	fmt.Println(s2, c2)

	s3 := "hello沙河"
	fmt.Println(len(s3))
	//遍历字符串

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i])
	}
	for k, v := range s3 {
		fmt.Printf("%d,%c\n", k, v)
	}

	//强制类型转换
	s5 := "big"
	//转换成字节数组
	betArr := []byte(s5)
	fmt.Println(betArr)
	betArr[0] = 'p'
	fmt.Println(betArr)
	s6 := string(betArr)
	fmt.Println(s6)
}
