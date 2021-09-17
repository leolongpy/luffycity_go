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

// 查询单条
func queryRowDemo() {
	var user User
	sqlstr := "select id,user_login,coin from cmf_user where id=1"
	err := DB.QueryRow(sqlstr).Scan(&user.id, &user.user_login, &user.coin)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	fmt.Printf("查询结果:%#v\n", user)
}

// 插入数据示例
func insertDemo() {
	sqlStr := "insert into cmf_user(user_login,coin) values (?,?)"
	name := "网络"
	coin := 10
	//Exec 执行
	ret, err := DB.Exec(sqlStr, name, coin)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsertid failed,err:%v\n", err)
		return
	}
	fmt.Println(theID)
}

// 查询多条数据
func queryMultiDemo() {
	var user User
	sqlStr := "select id,user_login,coin from cmf_user where id>?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	//循环读取数据
	for rows.Next() {
		err := rows.Scan(&user.id, &user.user_login, &user.coin)
		if err != nil {
			fmt.Printf("scan failed,err%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

// 更新数据
func updateDemo() {
	sqlStr := "update cmf_user set coin=? where id=?"
	ret, err := DB.Exec(sqlStr, 50, 2)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	// 拿到受影响的行数
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed,err%v\n", err)
		return
	}
	fmt.Println("受影响行数:", num)
}

//删除数据
func deleteDemo() {
	sqlStr := "delete from cmf_user where id=?"
	ret, err := DB.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
	}
	// 拿到受影响的行数
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed,err%v\n", err)
		return
	}
	fmt.Println("受影响行数:", num)
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/thinkcmf"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("connect mysql failed,err:%v\n", err)
	}
	//queryRowDemo()
	//insertDemo()
	//queryMultiDemo()
	//updateDemo()
	deleteDemo()
}
