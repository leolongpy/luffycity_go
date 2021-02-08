package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// 用户可以通过 -l 指定生成密码的长度
// 用户可以通过 -t 指定生成密码的字符集，
// 例如 -t num 生成全是数字的密码，-t char生成全是英文字符的密码，
// -t mix 生成包含数字和英文字符的密码，
// -t advance生成包含数字、英文以及特殊字符的密码

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

//解析方法
func parseArgs() {
	flag.IntVar(&length, "l", 10, "-l 是生成密码的长度参数")
	flag.StringVar(&charset, "t", "num", "-t 是置顶字符集 num是数字。。。")
	flag.Parse()
}

//生成密码的方法
func myPasswd() string {
	//存切片
	var passwd []byte = make([]byte, length, length)
	var res string
	if charset == "num" {
		res = NumStr
	} else if charset == "char" {
		res = CharStr
	} else if charset == "mix" {
		res = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		res = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		res = NumStr
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(res))
		passwd[i] = res[index]
	}
	return string(passwd)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	passwd := myPasswd()
	fmt.Println(passwd)
}
