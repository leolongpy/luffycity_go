package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		//写入数据
		err = createUser(username, pwd)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, "http://www.baidu.com", 301)
	} else {
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			w.WriteHeader(500)
		}
		err = t.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
		}

	}
}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		//去数据库校验
		err = queryUser(username, pwd)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, "http://www.baidu.com", 301)
	} else {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			w.WriteHeader(500)
		}
		err = t.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
		}
	}
}
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动http server失败！")
	}
}
