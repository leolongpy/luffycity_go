package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type user struct {
	Name  string
	Email string
	Phone int64
}

var (
	DefauleInterval = time.Minute * 1
	UserCache       = cache.New(DefauleInterval, DefauleInterval)
)

func GetUser(name string) user {
	res, found := UserCache.Get(name)
	if found {
		log.Printf("[在本地缓存中找到了对应的用户][name:%v][value:%v]", name, res.(user))
		return res.(user)
	} else {
		res := HttpGetUser(name)
		log.Printf("[本地缓存中没有找到对应的用户，去远端查询获取到了，塞入缓存中][name:%v][value:%v]", name, res)
		UserCache.Set(name, res, DefauleInterval)
		return res
	}
}
func HttpGetUser(name string) user {
	u := user{
		Name:  name,
		Email: "qq.com",
		Phone: time.Now().Unix(),
	}
	return u
}
func queryUser() {
	for i := 0; i < 10; i++ {
		userName := fmt.Sprintf("user_name_%d", i)
		GetUser(userName)
	}
}
func main() {
	log.Println("第一次query_user")
	queryUser()
	log.Println("第二次query_user")
	queryUser()
	queryUser()
	queryUser()
	queryUser()
	time.Sleep(61 * time.Second)
	log.Println("第三次query_user")
	queryUser()

}
