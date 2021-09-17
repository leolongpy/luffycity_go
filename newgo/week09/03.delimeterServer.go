package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()
		reader := bufio.NewReader(conn)
		for {
			data, err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}
			log.Println("recevied msg", len(data), "bytes:", string(data))
		}
	}
}
