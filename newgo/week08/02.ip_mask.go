package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid ip address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	_, bits := mask.Size()
	fmt.Println("地址是：", addr.String(),
		"Default mask length is ", bits,
		"Mask is (hex)", mask.String(),
		"Network is ", network.String(),
	)
	os.Exit(0)
}
