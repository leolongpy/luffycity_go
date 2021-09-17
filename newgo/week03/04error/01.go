package main

import (
	"errors"
	"log"
	"strings"
)

func validate(name string) (ok bool, err error) {
	if !strings.HasPrefix(name, "mysql") {
		return false, errors.New("name must start with mysql")
	}
	return true, nil
}

func main() {
	s1 := "mysql-abc"
	s2 := "redis-abc"
	_, err := validate(s1)
	if err != nil {
		log.Printf("[err:%v]", err)
	}
	if ok, err := validate(s2); err != nil {
		log.Printf("[err:%v][ok:%v]", err, ok)
	}
}
