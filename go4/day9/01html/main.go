package main

import (
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./demo.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		if username == "11" && pwd == "22" {
			t.Execute(w, "正确")
		} else {
			t.Execute(w, "错误")
		}
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
