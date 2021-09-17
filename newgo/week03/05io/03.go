package main

import (
	"io"
	"log"
	"os"
)

//type AlphaReader struct {
//	reader io.Reader
//}
//
//func (a *AlphaReader) Read(p []byte) (int ,error)  {
//	n,err := a.reader.Read(p)
//	if err != nil{
//		return n,err
//	}
//	buf := make([]byte,n)
//	for i:=0;i<n;i++ {
//		if char := guolv(p[i]);char!=0{
//			buf[i] = char
//		}
//	}
//	copy(p,buf)
//	return n,nil
//}
//func guolv(r byte) byte {
//	if (r>='A' && r<='Z') || (r>='a' && r<='z') {
//		return r
//	}
//	return 0
//}

func main() {
	file, err := os.Open("./01.go")
	if err != nil {
		log.Println(err)
		return
	}
	reader := AlphaReader{
		reader: file,
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
