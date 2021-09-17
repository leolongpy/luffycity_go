package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func queryUser(uid int) (User, error) {
	user := make(map[int]User)

	user[0] = User{"zs", 20}
	user[1] = User{"ls", 21}
	user[2] = User{"ww", 22}
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("%d err", uid)
}

func TestRPC(t *testing.T) {
	//编码中有一个字段是interface{}时，要注册一下
	gob.Register(User{})
	addr := "127.0.0.1:8000"
	// 创建服务端
	srv := NewServer(addr)
	srv.Register("queryUser", queryUser)
	//服务端等待调用
	go srv.Run()
	//客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
	}
	//创建客户端连接对象
	cli := NewClient(conn)
	var query func(int) (User, error)
	cli.callRPC("queryUser", &query)
	u, err := query(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u)
}
