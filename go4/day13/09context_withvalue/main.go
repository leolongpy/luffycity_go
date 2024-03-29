package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type TraceCode string
type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")

	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	useridKey := UserID("USER_ID")
	userid, ok := ctx.Value(useridKey).(int64)
	if !ok {
		fmt.Println("invalid user id")
	}
	log.Printf("%s worker func...", traceCode)
	log.Printf("userid:%d", userid)
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()

}

func main() {
	ctx,cancel := context.WithTimeout(context.Background(),time.Millisecond*50)
	ctx = context.WithValue(ctx,TraceCode("TRACE_CODE"),"123")
	ctx = context.WithValue(ctx,UserID("USER_ID"),int64(123456))
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
