package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	networkType := os.Args[1]
	service := os.Args[2]

	prot, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}
	fmt.Println("service port", prot)
	os.Exit(0)
}
