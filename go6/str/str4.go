package main

import "fmt"

func main() {
	str := "上海自来水来自海上"
	var r []rune = []rune(str)
	//反转
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str2 := string(r)
	if str2 == str {
		fmt.Println("是回文")
	}

}
