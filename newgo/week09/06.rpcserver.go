package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloServer struct {
}

func (p *HelloServer) Hello(req string, resp *string) error {
	*resp = "Hello" + req
	return nil
}

func main() {
	rpc.Register(new(HelloServer))
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal()
			continue
		}
		go func(conn net.Conn) {
			rpc.ServeConn(conn)
		}(conn)
	}
}
