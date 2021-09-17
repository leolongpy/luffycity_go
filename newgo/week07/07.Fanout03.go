package main

import (
	"fmt"
	"net"
	"time"
)

func handler(c net.Conn, ch chan int) {
	ch <- len(c.RemoteAddr().String())
	time.Sleep(10 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}
func pools(ch chan int, n int) {
	wch := make(chan int)

	for i := 0; i < n; i++ {
		go logger(wch)
	}
	for {
		wch <- <-ch
	}
}

func logger(ch chan int) {
	for {
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(<-ch)
	}
}

func server(l net.Listener, ch chan int) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}
func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan int)
	//go logger(ch)
	go pools(ch, 30)
	go server(l, ch)
	time.Sleep(10 * time.Second)

}
