package main

import (
	"fmt"
	"os"
)

var info = []map[string]string{}

func showMenu() {
	fmt.Println("1. 添加用户")
	fmt.Println("2. 查看用户")
	fmt.Println("0. 退出")
}

//修改数据
func updateinfo(name string) {
	for _, v := range info {
		if v["name"] == name {
			v["pwd"] = "666666"
			fmt.Println("密码已修改")
			return
		}
	}
	info = append(info, map[string]string{"name": name, "pwd": "666666"})
	fmt.Println("用户已添加")
	return
}

//获取数据
func getinfo() {
	if len(info) == 0 {
		fmt.Println("暂无数据")
	} else {
		for _, v := range info {
			fmt.Printf("姓名 %s,密码 %s\n", v["name"], v["pwd"])
		}
	}
	return
}
func main() {
	for {
		showMenu()
		// 获取用户输入的指令
		var i int
		fmt.Print("请输入指令：")
		fmt.Scanln(&i)
		fmt.Println("输入的选项是：", i)
		switch i {
		case 1:
			var name string
			fmt.Print("请输入姓名：")
			fmt.Scanln(&name)
			updateinfo(name)
		case 2:
			getinfo()
		case 0:
			// 退出
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}

}
