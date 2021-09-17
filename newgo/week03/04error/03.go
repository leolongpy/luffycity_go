package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}
func main() {
	err := errors.New("原始错误")
	myErr := MyError{
		err: err,
		msg: "数据上传问题",
	}
	fmt.Println(myErr.Error())
	e1 := fmt.Errorf("%w", err)
	fmt.Println(e1)
}
