package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func showMenu() {
	fmt.Println("1. 注册")
	fmt.Println("2. 登录")
	fmt.Println("0. 退出")
}

//注册
func register() {
	file, _ := os.OpenFile("./info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	var (
		name string
		pwd  string
		sex  string
	)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入密码：")
	fmt.Scanln(&pwd)
	for {
		fmt.Print("请输入性别：")
		fmt.Scanln(&sex)
		if sex != "男" && sex != "女" {
			fmt.Print("您输入性别有误，请重新输入")
		} else {
			break
		}
	}
	info := fmt.Sprintf("%s;%s;%s\n", name, pwd, sex)
	file.WriteString(info)
	fmt.Println("注册成功")
	return
}

//登录
func login() {
	file, err := os.Open("./info.txt")
	if err != nil {
		fmt.Println("暂无用户")
		return
	}
	defer file.Close()
	var (
		name string
		pwd  string
		sex  string
		flag bool
	)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入密码：")
	fmt.Scanln(&pwd)
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(lineBytes)
		lineSlice := strings.Split(lineStr, ";")
		if lineSlice[0] == name && lineSlice[1] == pwd {
			sex = lineSlice[2]
			flag = true
			break
		}
	}
	if flag {
		fmt.Println("异性列表")
		file.Close()
		file, err := os.Open("./info.txt")
		if err != nil {
			fmt.Println("暂无用户")
		}
		defer file.Close()
		reader2 := bufio.NewReader(file)
		for {
			lineBytes, _, err := reader2.ReadLine()
			if err == io.EOF {
				break
			}
			lineStr := string(lineBytes)
			lineSlice := strings.Split(lineStr, ";")
			if lineSlice[2] != sex {
				fmt.Println(lineSlice[0])
			}
		}
	} else {
		fmt.Println("您输入的姓名或密码错误")
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
			register()
		case 2:
			login()
		case 0:
			// 退出
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}

}
