package main

import (
	"fmt"
	"sync"
)

func send_num(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func recv_num(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(1)
	go send_num(ch, &wg)
	wg.Add(1)
	go recv_num(ch, &wg)
	wg.Wait()
}
