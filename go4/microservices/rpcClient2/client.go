package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// ArithRequest 参数结构体
type ArithRequest struct {
	A, B int
}

// ArithResponse 返回给客户端的结构体
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	coon, err := jsonrpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	req := ArithRequest{9, 2}
	var res ArithResponse
	err = coon.Call("Arith.Mul", req, &res)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	err = coon.Call("Arith.Div", req, &res)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.A, req.B, res.Quo, res.Rem)
}
