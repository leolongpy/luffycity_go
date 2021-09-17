package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// RPC计算圆的周长和面积

type Params struct {
	Width, Height int
}

type Rect struct {
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	//1.注册服务
	rect := new(Rect)
	rpc.Register(rect)
	//2.注册服务绑定到http协议上
	rpc.HandleHTTP()
	//3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
