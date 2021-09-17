package main

import (
	"io/ioutil"
	"log"
)

func main() {
	fs, _ := ioutil.ReadDir("../05io")
	for _, f := range fs {
		log.Printf("[name:%v][size:%v][mode:%v][modTime:%v]",
			f.Name(),
			f.Size(),
			f.Mode(),
			f.ModTime(),
		)
	}
}
