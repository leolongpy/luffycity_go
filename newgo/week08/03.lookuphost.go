package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := os.Args[1]
	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
