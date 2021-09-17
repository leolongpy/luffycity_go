package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 使用连接池方式连接mysql

var DB *sql.DB

func initDB(dsn string) (err error) {
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	//设置最大连接数
	DB.SetMaxOpenConns(50)
	//设置最大的空闲连接数
	//DB.SetMaxIdleConns(20)
	return nil
}

type User struct {
	id         int64
	user_login string
	coin       int64
}

// 预处理插入
func prepareInsertDemo() {
	sqlStr := "insert into cmf_user (user_login,coin) values (?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
	}
	defer stmt.Close()
	//执行重复插入命令
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("stu%02d", i)
		stmt.Exec(name, i)
	}
}

// 预处理查询
func prepareQueryDemo() {
	sqlStr := "select id,user_login,coin from cmf_user where id=?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare faild,err:%v\n", err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		rows, err := stmt.Query(i)
		if err != nil {
			fmt.Printf("query failed,err:%v\n", err)
			continue
		}
		defer rows.Close()
		var user User
		for rows.Next() {
			err := rows.Scan(&user.id, &user.user_login, &user.coin)
			if err != nil {
				fmt.Printf("scan failed,err:%v\n", err)
				return
			}
			fmt.Printf("user:%#v\n", user)
		}
	}
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/thinkcmf"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("connect mysql failed,err:%v\n", err)
	}
	//prepareInsertDemo()
	prepareQueryDemo()
}
