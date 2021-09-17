package main

//
//import (
//	"fmt"
//	"net"
//	"time"
//)
//
//func handler(c net.Conn,ch chan int)  {
//	ch<-len(c.RemoteAddr().String())
//	time.Sleep(10*time.Millisecond)
//	c.Write([]byte("ok"))
//	c.Close()
//}
//func logger(ch chan int)  {
//	for  {
//		time.Sleep(1500*time.Millisecond)
//		fmt.Println(<-ch)
//	}
//}
//
//func server(l net.Listener,ch chan int)  {
//	for  {
//		c,err := l.Accept()
//		if err != nil{
//			continue
//		}
//		go handler(c,ch)
//	}
//}
//func main() {
//	l,err :=net.Listen("tcp",":5000")
//	if err != nil {
//		panic(err)
//	}
//	ch := make(chan int)
//	go logger(ch)
//	go server(l,ch)
//	time.Sleep(10*time.Second)
//
//}
