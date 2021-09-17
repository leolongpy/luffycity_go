package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello沙河小王子"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

}
