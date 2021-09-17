package main

import (
	"io"
	"log"
	"net"
)

func Start() {
	linstener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer linstener.Close()

	for {
		conn, err := linstener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		for {
			var data = make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil && err != io.EOF {
				log.Println(err)
			}
			if n > 0 {
				log.Println("received msg", n, "bytes", string(data[:n]))
			}
		}
	}
}

func main() {
	Start()
}
