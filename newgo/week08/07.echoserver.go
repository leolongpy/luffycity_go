package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

func handleNewRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}

func main() {
	var host = flag.String("host", "", "host")
	var prot = flag.String("port", "3333", "port")
	flag.Parse()
	l, err := net.Listen("tcp", *host+":"+*prot)
	if err != nil {
		fmt.Println("Error listenning:", err.Error())
	}
	defer l.Close()
	fmt.Println("Listenning on " + *host + ":" + *prot)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accrpting:", err)
		}
		fmt.Printf("收到信息 %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleNewRequest(conn)

	}
}
