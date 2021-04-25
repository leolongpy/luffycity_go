package main

import (
	"fmt"
	"net"
)

func main() {
	//1.连接服务器
	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//写数据到服务端
	_, err = conn.Write([]byte("老铁"))
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 16)
	count, addr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("count", count)
	fmt.Println("addr", addr)
	fmt.Println("data:", string(data))

}
