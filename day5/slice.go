package main

import "fmt"

func main() {
	m1 := map[int]string{1: "zs", 2: "ls"}
	m1[1] = "ww"
	m1[3] = "zhaosi"
	fmt.Println(m1)
	//遍历
	for k, v := range m1 {
		fmt.Printf("%d----->%s\n", k, v)
	}

	for k := range m1 {
		fmt.Printf("%d---->%s\n", k, m1[k])
	}

	value,ok :=m1[5]
	fmt.Println("value=",value,",ok=",ok)

	delete(m1,1)
	fmt.Println(m1)
}
