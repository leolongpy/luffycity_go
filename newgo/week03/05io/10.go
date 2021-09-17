package main

import (
	"fmt"
	"os"
)

//os.Stdout.Write 替代fmt.print
func main() {
	fmt.Println("123")
	os.Stdout.Write([]byte("456"))
}
