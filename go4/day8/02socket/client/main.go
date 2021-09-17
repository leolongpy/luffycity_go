package main

import (
	"fmt"
	"luffycity_go/go4/day8/02socket/proto"
	"net"
)

// 粘包现象 客户端
// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dia; failed,err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello,hello.How are you`
		//调用自己定义的协议
		pkg, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:", err)
			return
		}
		conn.Write(pkg)
	}
}
