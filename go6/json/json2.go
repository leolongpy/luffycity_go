package main

import (
	"encoding/json"
	"fmt"
)

//map 生成 json
func main() {
	mmp := make(map[string]interface{})
	mmp["name"] = "Mr.Sun"
	mmp["age"] = 18
	mmp["niubility"] = true
	mjson, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(mjson))
}
