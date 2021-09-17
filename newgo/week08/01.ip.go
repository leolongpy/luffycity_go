package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid ip adderss")
	} else {
		fmt.Println("the address is", addr.String())
	}
	os.Exit(0)
}
