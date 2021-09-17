package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

//将表抽象为对象
type KfPerson struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Idcard string `db:"idcard"`
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println("ERROR!", err, why)
	}
}

//定义全局变量
var (
	//读写数据的管道
	chanData chan *KfPerson
	db       *sqlx.DB
	wg       sync.WaitGroup
)

//判断标记
func CheckIfFileExist(filenme string) (exists bool) {
	//判断文件是否存在
	fileInfo, err := os.Stat(filenme)
	if fileInfo != nil && err == nil {
		exists = true
	} else {
		exists = false
	}
	return
}

//导入数据
func init() {
	exists := CheckIfFileExist("E:/Desktop/kaifang_good_done.mark")
	if exists {
		fmt.Println("数据已初始化")
		return
	}
	//连接数据库
	var err error
	db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/thinkcmf")
	HandleError(err, "sql open")
	defer db.Close()
	fmt.Println("数据库已连接")
	_, err = db.Exec("create table if not exists kfperson (id int primary key auto_increment,name varchar(20),idcard char(18))")
	HandleError(err, "create table")
	fmt.Println("表已创建")
	//初始化管道
	chanData = make(chan *KfPerson, 1000000)
	//打开100协程插入数据
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go insertKfPerson()
	}
	//读取数据
	file, e := os.Open("E:/Desktop/kfang_good.txt")
	HandleError(e, "os.open")
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println("数据文本已打开")
	for {
		lineBytes, _, err := reader.ReadLine()
		HandleError(err, "reader.readLine")
		if err == io.EOF {
			close(chanData)
			break
		}
		lineStr := string(lineBytes)
		//切分
		fields := strings.Split(lineStr, ",")
		name, idcard := fields[0], fields[1]
		//区空格
		name = strings.TrimSpace(name)
		if len(name) > 20 {
			fmt.Println("名字过长不要了")
			continue
		}
		kfPerson := KfPerson{Name: name, Idcard: idcard}
		//导入管道
		chanData <- &kfPerson
	}
	fmt.Println("数据初始化成功")
	//创建标记
	_, err = os.Create("E:/Desktop/kaifang_good_done.mark")
	HandleError(err, "os.Crate")
	wg.Wait()
}

//将数据管道数据插入到数据库中
func insertKfPerson() {
	//遍历管道，拿数据
	for KfPerson := range chanData {
		for {
			result, err := db.Exec("insert into kfperson(name,idcard) values (?,?)", KfPerson.Name, KfPerson.Idcard)
			HandleError(err, "db insert")
			if err != nil {
				<-time.After(5 * time.Second)
			} else {
				if n, e := result.RowsAffected(); e == nil && n > 0 {
					fmt.Printf("插入%s成功\n", KfPerson.Name)
					break
				}
			}
		}
	}
	wg.Done()

}

//抽象缓冲结构体
type QueryResult struct {
	value     []KfPerson
	cacheTime int64
	count     int
}

// 获取加入缓存的时间
func (qr *QueryResult) GetCachTime() int64 {
	return qr.cacheTime
}

//实现缓存策略，删除最早加入的
func UpdateCache(cacheMap *map[string]QueryResult) (delKey string) {
	myTime := time.Now().UnixNano()
	for key, timeData := range *cacheMap {
		if timeData.GetCachTime() < myTime {
			myTime = timeData.GetCachTime()
			delKey = key
		}
	}
	delete(*cacheMap, delKey)
	return delKey
}
func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/thinkcmf")
	HandleError(err, "sql open")
	defer db.Close()
	kfMap := make(map[string]QueryResult, 0)

	for {
		var name string
		fmt.Print("请输入要查询的姓名：")
		fmt.Scanf("%s", &name)
		if name == "exit" {
			break
		}
		if name == "cache" {
			fmt.Printf("共缓存了%d条结果:\n", len(kfMap))
			for key := range kfMap {
				fmt.Println(key)
			}
			continue
		}
		//操作缓存查询
		if qr, ok := kfMap[name]; ok {
			qr.count += 1
			fmt.Printf("查询到%d条结果：\n", len(qr.value))
			fmt.Println(qr.value)
			continue
		}
		//没有缓存，走数据库
		kfpeople := make([]KfPerson, 0)
		e := db.Select(&kfpeople, "select id,`name`,idcard from kfperson where name=?", name)
		fmt.Printf("查询到%d条结果:\n", len(kfpeople))
		HandleError(e, "select name")
		fmt.Println(kfpeople)
		queryResult := QueryResult{value: kfpeople}
		queryResult.cacheTime = time.Now().UnixNano()
		queryResult.count = 1
		kfMap[name] = queryResult
		if len(kfMap) > 2 {
			delKey := UpdateCache(&kfMap)
			fmt.Printf("%s已经被淘汰处缓存\n", delKey)
		}

	}
	fmt.Println("done")
}
