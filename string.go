package main

import (
	"fmt"
	"strings"
)

func main()  {
	s1:="alexsb"
	fmt.Println(len(s1))
	s2:="python"
	fmt.Println(s1+s2)
	s3:= fmt.Sprintf("%s-------%s",s1,s2)
	fmt.Println(s3)
	// 分割
	ret := strings.Split(s1,"x")
	fmt.Println(ret)
	//判断是否包含
	ret2 := strings.Contains(s1,"sb")
	fmt.Println(ret2)
	//判断前缀和后缀
	ret3:=strings.HasPrefix(s1,"alex")
	ret4:=strings.HasSuffix(s1,"sb")
	fmt.Println(ret3,ret4)
	//求子串的位置
	s4:="applepen"
	fmt.Println(strings.Index(s4,"p"))
	fmt.Println(strings.LastIndex(s4,"p"))
	// join
	a1 := []string{"python","php","java","golang"}
	fmt.Println(strings.Join(a1,"-"))
}

