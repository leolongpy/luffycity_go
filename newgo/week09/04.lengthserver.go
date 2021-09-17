package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
			peek, err := reader.Peek(4)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
					continue
				} else {
					break
				}
			}

			buffer := bytes.NewReader(peek)
			var length int32
			err = binary.Read(buffer, binary.BigEndian, &length)
			if err != nil {
				log.Println(err)
			}
			if int32(reader.Buffered()) < length+4 {
				continue
			}
			data := make([]byte, length+4)
			_, err = reader.Read(data)
			if err != nil {
				continue
			}
			log.Println("received msg", string(data[4:]))
		}
	}
}
