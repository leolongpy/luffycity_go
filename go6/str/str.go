package main

import (
	"fmt"
	"strings"
)

func main() {
	//str1 := "hello"
	//fmt.Println(len(str1))
	//
	//str2 := "hello你好"
	//fmt.Println(len(str2))
	//
	//var r []rune = []rune(str2)
	//fmt.Println(r)
	//fmt.Println(len(r))
	//获取字符串长度
	fmt.Println(len("hello"))
	//字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("hello", "llo"))
	//判断字符串s是否以prefix为开头
	fmt.Println(strings.HasPrefix("hello", "he"))
	fmt.Println(strings.HasSuffix("hello", "lo"))
	//字符串连接
	s := []string{"abc", "456", "999"}
	fmt.Println(strings.Join(s, "** "))
	//在字符串s中查找sep所在的位置，返回位置值，找不到返回-1  LastIndex是从后往前查找
	fmt.Println(strings.Index("chicken", "ken"))
	// 重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("ok ok ok", "k", "ky", 2))
	// 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	// 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]\n", strings.Trim(" !哈!哈! ", "! "))
	// 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields(" a b c  "))
}
