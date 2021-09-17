package main

import (
	"io"
	"log"
	"strings"
)

type AlphaReader struct {
	reader io.Reader
}

func (a *AlphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := guolv(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}
func guolv(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func main() {
	originReader := strings.NewReader("sdf  #%dsfg ds  5465")
	reader := AlphaReader{
		reader: originReader,
	}
	p1 := make([]byte, 4)
	for {
		n1, err := reader.Read(p1)
		if err == io.EOF {
			break
		}
		log.Printf("[内容:%v]", string(p1[:n1]))
	}
}
