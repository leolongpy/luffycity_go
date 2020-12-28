package main

import (
	"fmt"
	"os"
)

//定义结构体
type book struct {
	title   string
	author  string
	price   float32
	publish bool
}

//定义创建新书的构造函数
func newBook(title, author string, price float32, publish bool) *book {
	return &book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

//定义一个存放book指针的切片，用来存储所有的书籍
var allBooks = make([]*book, 0, 200)

//打印菜单
func showMenu() {
	fmt.Println("欢迎登录BMS!")
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
}

// 定义一个专门用来获取用户输入的书籍信息的
func userInput() *book {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名：")
	fmt.Scanln(&title)
	fmt.Print("请输入作者:")
	fmt.Scanln(&author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架[true|false]:")
	fmt.Scanln(&publish)
	fmt.Println(title, author, price, publish)
	book := newBook(title, author, price, publish)
	return book
}

//添加书籍的函数
func addBook() {
	book := userInput()
	for _, b := range allBooks {
		if b.title == book.title {
			fmt.Printf("《%s》这本书已经存在了！", book.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Println("书籍添加成功")
}

//修改书籍函数
func updateBook() {
	book := userInput()
	for index, b := range allBooks {
		if b.title == book.title {
			allBooks[index] = book
			fmt.Printf("书名:《%s》更新成功！\n", book.title)
			return
		}
	}
	fmt.Printf("书名:《%s》不存在！\n", book.title)
}

//展示书籍函数
func showBooks() {
	if len(allBooks) == 0 {
		fmt.Println("啥也没有")
		return
	}
	for _, b := range allBooks {
		fmt.Printf("《%s》 作者：%s 价格：%.2f 是否上架销售：%t\n", b.title, b.author, b.price, b.publish)
	}
}
func main() {
	for {
		showMenu()
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addBook()
		case 2:
			updateBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)
		}
	}
}
