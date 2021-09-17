package main

import "log"

type notifer interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	log.Printf("[普通用户的通知][notify to user:%s]", u.name)
}

type admin struct {
	name string
	age  int
}

func (a *admin) notify() {
	log.Printf("[管理员的通知][notify to user:%s]", a.name)
}

//多态的统一调用入口
func sendNotify(n notifer) {
	n.notify()
}

func main() {
	u1 := user{
		name:  "leo",
		email: "12@qq.com",
	}
	a1 := admin{
		name: "long",
		age:  18,
	}
	log.Println("直接调用结构体绑定的方法")
	u1.notify()
	a1.notify()
	log.Println("体现多态")
	sendNotify(&u1)
	sendNotify(&a1)
	log.Println("多态承载容器")
	ns := make([]notifer, 0)
	ns = append(ns, &a1)
	ns = append(ns, &u1)
	for _, n := range ns {
		n.notify()
	}
}
