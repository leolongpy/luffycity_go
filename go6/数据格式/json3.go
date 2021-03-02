package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility"`
}

func main() {
	b := []byte(`{"age":"18","name":"Mr.Sun","niubility":true}`)
	var p Person2
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
