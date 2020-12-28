package main

import "fmt"

func main() {
	//var a int
	//fmt.Println(a)
	//b := &a
	//fmt.Printf("b=%v\n", b)
	//fmt.Printf("type b :%T\n", b)
	//
	//c := "豪杰"
	//fmt.Printf("&V=%v\n", &c)
	//
	//d := 100
	//b = &d
	//fmt.Println(b)
	//fmt.Println(*b)
	//fmt.Println(b == &d)
	a := [3]int{1, 2, 3}
	modifyArray(a)
	fmt.Println(a)

	modfilyArray2(&a)
	fmt.Println(a)

}

func modifyArray(a1 [3]int) {
	a1[0] = 100
}

func modfilyArray2(a1 *[3]int) {
	//(*a1)[0]= 100
	a1[0] = 100
}
