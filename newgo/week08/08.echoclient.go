package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello" + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of", e.Error())
			break
		}
	}
}
func handeRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of", err.Error())
	}
	fmt.Println(string(buf[:n-1]))
}

func main() {
	var host = flag.String("host", "", "host")
	var port = flag.String("port", "3333", "port")

	flag.Parse()

	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error Connecting：", err.Error())
	}
	defer conn.Close()

	fmt.Println("connecting to ", *host+":"+*port)
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handeRead(conn, &wg)
	wg.Wait()
}
