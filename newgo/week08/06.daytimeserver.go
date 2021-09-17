package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func main() {
	srevice := ":12000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", srevice)
	checkError(err)

	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listner.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		daytime := time.Now().String()
		println(daytime)
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
