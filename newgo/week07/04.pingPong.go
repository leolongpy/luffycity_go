package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player("1", table)
	go player("2", table)
	table <- Ball
	time.Sleep(1 * time.Second)

}

func player(id string, table chan int) {
	for {
		ball := <-table
		fmt.Printf("%s got ball [%d]\n", id, ball)
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("%s bonceback ball [%d]\n", id, ball)
		ball++
		table <- ball
	}
}
