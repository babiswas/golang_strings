package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 3
	ch <- 3
	close(ch)
	go sum(ch)
	time.Sleep(100 * time.Millisecond)
}

func sum(ch <-chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Println(sum)
}
