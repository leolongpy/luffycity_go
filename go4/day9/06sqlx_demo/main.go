package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int64  `db:"id"`
	UserLogin string `db:"user_login"`
	Coin      int64  `db:"coin"`
}

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/thinkcmf"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return nil
}

// 查询单条数据
func queryRowDemo() {
	sqlStr := "select id,user_login from cmf_user where id=?"
	var user User
	err := DB.Get(&user, sqlStr, 1)
	if err != nil {
		fmt.Println(err, 1)
		return
	}
	fmt.Printf("user:%#v\n", user)
}

// 查询多行
func queryMultidemo() {
	sqlStr := "select id,user_login,coin from cmf_user where id>?"
	var users []*User
	err := DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}

}

//事务操作
func transDemo() {
	tx, err := DB.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println(err)
		return
	}
	sql1 := "update cmf_user set coin=coin+? where id=?"
	tx.MustExec(sql1, 2, 1)

	sql2 := "update cmf_user set coin=coin-? where id=?"
	tx.MustExec(sql2, 2, 2)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	fmt.Println("两条数据程序成功")

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect mysql failed,err:%v\n", err)
	}
	//queryRowDemo()
	//queryMultidemo()
	transDemo()
}
