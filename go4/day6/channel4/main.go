package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var itemChan chan *item
var resultChan chan *result
var wg sync.WaitGroup

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

//生产者
func producer(ch chan *item) {
	//产生随机数
	var id int64
	for i := 0; i < 10000; i++ {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}
		//把随机数发送到管道中
		ch <- tmp
	}
	close(ch)
}

//计算一个数字每个位的和
func calc(num int64) int64 {
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

//消费者
func comsumer(ch chan *item, resultChan chan *result) {
	defer wg.Done()
	for tmp := range ch {
		sum := calc(tmp.num)
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}
}
func startWoker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go comsumer(ch, resultChan)
	}
}

//打印结果
func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id:%v,num:%v,sum%v\n", ret.item.id, ret.item.num, ret.sum)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	itemChan = make(chan *item, 10000)
	resultChan = make(chan *result, 10000)
	go producer(itemChan)
	wg.Add(20)
	startWoker(20, itemChan, resultChan)
	wg.Wait()
	close(resultChan)
	printResult(resultChan)

}
