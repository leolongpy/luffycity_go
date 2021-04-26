package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"_"`
	Hobby string `json:"hobby"`
}

func main() {
	p := Person{"zs", "女"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
}
