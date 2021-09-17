package main

import "os"

func main() {
	file, _ := os.Create("b.txt")

	for i := 0; i < 5; i++ {
		file.WriteString("WriteString\n")
		file.Write([]byte("Write\n"))
	}
}
