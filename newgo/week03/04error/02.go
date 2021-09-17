package main

import (
	"log"
	"os"
)

func main() {
	f1, err := os.Stat("text.txt")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			log.Println("PathError")
		case *os.LinkError:
			log.Println("LinkError")
		case *os.SyscallError:
			log.Println("SyscallError")

		}
	} else {
		log.Printf("f1:%v", f1)
	}
}
