package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"age":1.3,"name":"Mr.Sun","niubility":true}`)
	//声明接口
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}
	}
}
