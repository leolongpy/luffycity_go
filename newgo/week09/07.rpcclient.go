package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	var reply string

	err = client.Call("HelloServer.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
