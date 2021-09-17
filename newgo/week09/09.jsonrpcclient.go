package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	var reply int
	type Args struct {
		A, B int
	}
	arges := &Args{7, 8}
	err = client.Call("Arith.Add", arges, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
