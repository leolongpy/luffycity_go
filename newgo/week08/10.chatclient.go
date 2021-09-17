package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func SendMsg(conn net.Conn) {
	username := conn.LocalAddr().String()
	for {
		var input string
		fmt.Scanln(&input)

		if input == "/q" || input == "/quite" {
			fmt.Println("Byebye...")
			conn.Close()
			os.Exit(0)
		}
		// 处理消息
		if len(input) > 0 {
			msg := username + " say " + input
			_, err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}

func StartClient(tcpAddrStr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil {
		log.Printf("Resovle tcp addr failed: %v\n", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Dial tcp addr failed: %v\n", err)
		return
	}

	buf := make([]byte, 1024)

	go SendMsg(conn)

	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}
		fmt.Println(string(buf[:length-1]))
	}
}

func main() {
	StartClient(os.Args[1])
}
