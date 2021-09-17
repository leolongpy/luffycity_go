package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("err:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolnesd address is ", addr.String())
}
