package main

import (
	"fmt"
	"log"
	"net"
)

// BoadMessage 向所有的群友发广播
func BoadMessage(conns *map[string]net.Conn, message chan string) {
	for {
		msg := <-message
		fmt.Println(msg)

		for key, conn := range *conns {
			fmt.Println("connection is connected from ", key)
			_, err := conn.Write([]byte(msg))
			if err != nil {
				log.Printf("broad message to %s failed:%v\n", key, err.Error())
				delete(*conns, key)
			}
		}
	}
}

// Start 启动
func Start(port string) {
	host := ":" + port

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Printf("resovle tcp addr failed: %v\n", err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("listen tcp addr failed: %v\n", err)
		return
	}
	// 建立连接池
	conns := make(map[string]net.Conn)
	//消息通道
	messageChange := make(chan string, 10)
	//广播消息
	go BoadMessage(&conns, messageChange)
	for {
		fmt.Printf("listening port %s ...\n", port)
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("接收失败:%v", err)
			continue
		}
		conns[conn.RemoteAddr().String()] = conn
		fmt.Println(conns)

		//处理消息
		go Handler(conn, &conns, messageChange)
	}
}

// Handler 处理客户端发送到服务端的消息
func Handler(conn net.Conn, conns *map[string]net.Conn, message chan string) {
	fmt.Println("connt from client ", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("read client message failed: %v\n", err)
			delete(*conns, conn.RemoteAddr().String())
			conn.Close()
			break
		}
		recvStr := string(buf[:length-1])
		message <- recvStr
	}
}
func main() {
	port := "9090"
	Start(port)
}
