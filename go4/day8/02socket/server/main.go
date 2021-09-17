package main

import (
	"bufio"
	"bytes"
	"fmt"
	"luffycity_go/go4/day8/02socket/proto"
	"net"
)

// 粘包现象 服务端
// socket_stick/server/main.go

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go procsee(conn)
	}

}

func procsee(conn net.Conn) {
	defer conn.Close()
	//reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	//循环度
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		reader := bufio.NewReader(bytes.NewReader(buf[:n]))
		msg, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("decode mag failed,err:", err)
			return
		}
		fmt.Println("收到客户端发来的数据", msg)
	}
}
