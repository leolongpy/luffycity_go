package main

type Book struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
	Price string `db:"price"`
}
