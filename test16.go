package main

import (
	"fmt"
	"sync"
)

func process_odd(odd chan<- int, evn <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 100; {
		odd <- i
		evn_num, ok := <-evn
		if ok == false {
			break
		}
		fmt.Println(evn_num)
		i = i + 2
	}
	close(odd)
	evn_num, ok := <-evn
	if ok == true {
		fmt.Println(evn_num)
	}
}

func process_evn(odd <-chan int, evn chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i < 100; {
		odd_num, ok := <-odd
		if ok == false {
			break
		}
		fmt.Println(odd_num)
		evn <- i
		i = i + 2
	}
	close(evn)
	odd_num, ok := <-odd
	if ok == true {
		fmt.Println(odd_num)
	}
}

func main() {
	odd := make(chan int)
	evn := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go process_odd(odd, evn, &wg)
	wg.Add(1)
	go process_evn(odd, evn, &wg)
	wg.Wait()
}
