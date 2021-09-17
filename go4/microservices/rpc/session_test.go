package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWrite(t *testing.T) {
	addr := "127.0.0.1:8000"
	my_data := "龙海晨"
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err.Error())
		}
		conn, _ := lis.Accept()
		s := Session{conn: conn}
		err = s.Write([]byte(my_data))
		if err != nil {
			t.Fatal(err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err.Error())
		}
		s := Session{conn: conn}
		data, err := s.Read()
		if err != nil {
			t.Fatal(err.Error())
		}
		if string(data) != my_data {
			t.Fatal(err.Error())
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
