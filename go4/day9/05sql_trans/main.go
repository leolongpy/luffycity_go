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

//事务
func transDemo() {
	tx, err := DB.Begin() //开启事务
	if err != nil {
		fmt.Println(err)
		if tx != nil {
			tx.Rollback()
		}
		return
	}
	sql1 := "update cmf_user set coin=coin+? where id=?"
	_, err = tx.Exec(sql1, 2, 1)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}
	sql2 := "update cmf_user set coin=coin-? where id=?"
	_, err = tx.Exec(sql2, 2, 2)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}
	fmt.Println("两行记录已更新")
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/thinkcmf"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("connect mysql failed,err:%v\n", err)
	}
	transDemo()

}
