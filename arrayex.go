package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 3, 5, 7}
	sum := 0
	for _, v := range a1 {
		fmt.Println(v)
		sum = sum + v
	}
	fmt.Println(sum)

	// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	//   遍历数组，
	// 2.1依次取出每个元素
	// 2.2 计算一下 other= 8-当前值
	// 2.3 在不在数组中，在的话把索引拿出来
	for k,v:=range a1{
		o:=8-v
		for k2:=k+1;k2<len(a1);k2++ {
			if a1[k2]==o{
				fmt.Printf("它们的索引是：(%d %d)\n", k, k2)
			}
		}
	}
}
