package main

import "fmt"

func main() {
	var  (
		name string
		age int
		married bool
	)
	//fmt.Scan(&name,&age,&married)
	fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
	//fmt.Scanln(&name, &age, &married)
	fmt.Println(name,age,married)

}
