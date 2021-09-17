package main

import "fmt"

func main() {
	//var a = [3]int{1, 2, 3}
	////直接声明切片
	//var b = []int{1, 2, 3}
	//fmt.Println(a, b)
	//fmt.Printf("a:%T  b:%T\n", a, b)
	//fmt.Printf("b的长度是：%d\n", len(b))
	//fmt.Printf("b的容量是：%d\n", cap(b))
	//fmt.Println(b[1])
	////从数组中得到切片
	//var c []int
	//c = a[:]
	//fmt.Printf("c:%T\n", c)
	//fmt.Println(c)
	//
	//d := a[:2]
	//e := a[1:]
	//fmt.Println(d)
	//fmt.Println(e)
	//
	////切片大小（目前元素的数量）
	//fmt.Println(len(b))
	////容量（底层数组最大能放多少元素）
	//x:=[...]string{"北京", "上海", "深圳", "广州", "成都", "杭州", "重庆"}
	//y:=x[1:4]
	//fmt.Println(y)
	//fmt.Printf("切片y的长度是：%d\n", len(y))
	//fmt.Printf("切片y的容量是：%d\n", cap(y))

	//var a = []int{} //空切片
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1) // [1 1 1]  3  4
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	//a = append(a, 1)
	//fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)

	////定义一个数组
	//a := [...]int{1, 3, 5, 7, 9,11, 13}
	//b:=a[:]
	//b[0] = 100
	//fmt.Println(a[0])
	//fmt.Printf("b:%p\n", b)
	//
	//c:=a[2:5]
	//fmt.Println(c)
	//fmt.Println(len(c))//3
	//fmt.Println(cap(c))//5
	//fmt.Printf("c:%p\n",c)
	//
	//d:=c[:5]
	//fmt.Println(d)
	//fmt.Println(len(d)) //5
	//fmt.Println(cap(d)) //5
	//fmt.Printf("d:%p\n", d)
	//
	//e := d[2:]
	//fmt.Println(e)      //[9 11 13]
	//fmt.Println(len(e)) //3
	//fmt.Println(cap(e)) //3
	//fmt.Printf("e:%p\n", e)
	//
	//e = append(e,100,200)
	//fmt.Println(e)
	//fmt.Println(len(e)) //5
	//fmt.Println(cap(e)) //6
	//fmt.Printf("e:%p\n", e)
	//
	//e[0] = 900
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(e)
	//
	//m:=[]int{1,2,3,4}
	//fmt.Println(m)
	//fmt.Printf("%p\n",m)
	//m=append(m[:1],m[3:]...)
	//fmt.Printf("%p\n cap:%d  len:%d\n", m, cap(m), len(m))
	//fmt.Println(m)

	a := [...]int{1}
	b := a[:]
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	b[0] = 100
	fmt.Println(a, b)
	//扩容策略
	fmt.Println("b的容量：", cap(b))
	b = append(b, 3, 4, 5, 6, 7, 8)
	fmt.Println("b的容量：", cap(b))
	b = append(b, 8)
	fmt.Println("b的容量：", cap(b))

	var s1 []int64
	s1 = append(s1, 100)
	fmt.Println(s1)
	s1 = make([]int64, 3)
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]string, 3, 10)
	fmt.Println(s2, len(s2), cap(s2))

}
