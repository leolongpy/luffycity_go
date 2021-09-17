package main

import "log"

type Person4 struct {
	Name     string
	Age      int
	Tags     map[string]string
	HouseID1 [2]int //数组是值类型
	HouseID2 []int  //切片是引用类型
}

func main() {
	p1 := Person4{
		Name:     "leo",
		Age:      123,
		Tags:     map[string]string{"k1": "v1", "k2": "v2"},
		HouseID1: [2]int{100, 102},
		HouseID2: []int{2000, 201},
	}
	p2 := p1

	//修改两个值类型的字段
	p1.Age = 19
	p2.Name = "long"
	//修改map
	p1.Tags["k1"] = "v11"
	//修改array
	p2.HouseID1[0] = 300
	//修改切片
	p1.HouseID2[1] = 301
	log.Printf("[p1的内存地址:%p][value:%+v]", &p1, p1)
	log.Printf("[p2的内存地址:%p][value:%+v]", &p2, p2)
}
