package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="hello hello how are you i am fine thank you"
	wordSlice:=strings.Split(s," ")
	fmt.Println(wordSlice)

	wordMap := make(map[string]int,len(wordSlice))
	for _,word:=range wordSlice{
		v,ok:=wordMap[word]
		if ok {
			wordMap[word] = v+1
		}else {
			wordMap[word] = 1
		}
	}
	for k,v := range wordMap{
		fmt.Println(k,v)
	}
}
