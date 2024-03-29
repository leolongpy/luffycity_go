package main

import (
	"fmt"
	"net"
)

func main() {
	udp := "0.0.0.0:8888"
	udpAddr, _ := net.ResolveUDPAddr("udp4", udp)
	//UDP的服务监听
	listen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()
	for {
		//缓冲区
		var data [1024]byte
		//接收 UDP 传输
		count, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("data:%s addr:%v \n", string(data[0:count]), addr)
		//返回信息
		_, err = listen.WriteToUDP([]byte("666"), addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
