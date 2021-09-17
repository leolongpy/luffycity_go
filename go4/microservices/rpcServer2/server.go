package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Arith 用于注册的结构体
type Arith struct {
}

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

func (a *Arith) Mul(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (a *Arith) Div(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("a new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
