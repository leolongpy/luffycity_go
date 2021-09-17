package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/thinkcmf"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed,err:%v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect Mysql failed,err:%v\n", err)
		return
	}
	fmt.Println("数据库链接成功")
}
