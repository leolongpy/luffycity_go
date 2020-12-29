package main

import "fmt"

type xiyiji interface {
	wash()
	dry()
}
type Haier struct {
	name  string
	price float64
	mode  string
}
type tianluo struct {
	name string
}

func (t tianluo) wash() {
	fmt.Println("田螺姑娘可以洗衣服~")
}
func (t tianluo) dry() {
	fmt.Println("田螺姑娘可以把衣服拧干~")
}

func (h Haier) wash() {
	fmt.Println("海尔洗衣机能洗衣服~")
}
func (h Haier) dry() {
	fmt.Println("海尔洗衣机自带甩干~")
}

func main() {
	var a xiyiji
	h1 := Haier{
		name:  "小神童",
		price: 998.98,
		mode:  "滚筒",
	}
	fmt.Printf("%T\n", h1)
	a = h1
	fmt.Println(a)
	tl := tianluo{
		name: "螺蛳粉",
	}
	a = tl
	fmt.Println(a)
}
