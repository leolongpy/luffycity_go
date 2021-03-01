package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := make([]byte, 0, 10)
	//以10进制方式追加
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcd")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}
