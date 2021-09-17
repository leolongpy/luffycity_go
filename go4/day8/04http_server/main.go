package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./form.html")
	if err != nil {
		fmt.Println("read from file failed,err:", err)
		return
	}
	w.Write(data)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求的方法", r.Method)
	r.ParseForm() //解析
	fmt.Printf("%#v\n", r.Form)
	userNameValue := r.Form.Get("username")
	pwdValue := r.Form.Get("pwd")
	fmt.Println(userNameValue, pwdValue)
	w.Write([]byte("index"))
}

func main() {
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
