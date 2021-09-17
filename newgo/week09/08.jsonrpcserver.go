package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}
type Arith int

func (a *Arith) Add(args *Args, resp *int) error {
	*resp = args.A * args.B
	return nil
}

func main() {
	rpc.Register(new(Arith))
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
