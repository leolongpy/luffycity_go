package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	tcp := "0.0.0.0:8888"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", tcp)
	//创建tcp服务端监听
	listenner, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listenner.Close()
	//服务端不断等待请求处理
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ClientConn(conn)
	}
}

// ClientConn 处理服务端逻辑
func ClientConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的ip地址
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")
	buf := make([]byte, 1024)
	for {
		//n是读取的长度
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		//切出有效数据
		result := buf[:n]
		fmt.Printf("接收到数据，来自[%s]  [%d]:%s\n", ipAddr, n, string(result))
		if string(result) == "exit" {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		//回复客户端
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}
