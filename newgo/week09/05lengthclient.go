package main

import (
	"bytes"
	"encoding/binary"
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
		data, err := Encode(strconv.Itoa(i) + "[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城]")
		_, err = conn.Write(data)
		data, err = Encode(strconv.Itoa(i) + "aaaa")
		_, err = conn.Write(data)
		data, err = Encode(strconv.Itoa(i) + "bbbb")
		_, err = conn.Write(data)
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second)
}

func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
