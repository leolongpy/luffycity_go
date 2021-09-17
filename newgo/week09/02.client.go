package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte(strconv.Itoa(i) + "[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城]\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "bbbb\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "cccc\n"))
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second)
}
