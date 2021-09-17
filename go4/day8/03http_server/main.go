package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP server端
func sayHello(w http.ResponseWriter, t *http.Request) {
	//fmt.Fprintln(w,"hello")
	//w.Write([]byte("hello"))
	data, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println("read from file failed,err:", err)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/", sayHello)
	//启动服务
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		panic(err)
	}
}
