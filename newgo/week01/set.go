package main

import (
	"fmt"
	"sync"
)

var Lock sync.Mutex

//slice去重
func removeSilce(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

//初始化
func NewSet(set []string) []string {
	Lock.Lock()
	defer Lock.Unlock()
	newSet := removeSilce(set)
	return newSet
}

//添加元素
func Add(str string, set []string) []string {
	Lock.Lock()
	defer Lock.Unlock()
	set = append(set, str)
	newSet := removeSilce(set)
	return newSet
}

//删除元素
func Del(str string, set []string) []string {
	Lock.Lock()
	defer Lock.Unlock()
	for k, v := range set {
		if v == str {
			set = append(set[:k], set[k+1:]...)
			break
		}
	}
	return set
}

//合并set
func Merge(set, newSet []string) []string {
	Lock.Lock()
	defer Lock.Unlock()
	set = append(set, newSet...)
	set = removeSilce(set)
	return set
}

//打印元素
func PrintElement(set []string) {
	for _, v := range set {
		fmt.Println(v)
	}
}
func JudgeElement(str string, set []string) bool {
	for _, v := range set {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	var set = []string{"a", "b", "b"}
	set = NewSet(set)
	set = Add("c", set)
	set = Del("a", set)
	set = Merge(set, []string{"d", "r", "f"})
	fmt.Println(set)
	PrintElement(set)
	fmt.Println(JudgeElement("r", set))
}
