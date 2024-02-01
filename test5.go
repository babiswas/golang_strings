package main

import (
	"fmt"
	"sync"
)

func num_generator(odd chan<- int, even chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func proc_ev_odd(even <-chan int, odd <-chan int, collector chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var isodd bool
	var iseven bool
	isodd = false
	iseven = false
	fmt.Println(<-even)
	iseven = true

	for (isodd) || (iseven) {
		if iseven == true {
			num, ok := <-odd
			if ok == false {
				break
			}
			collector <- num
			isodd = true
			iseven = false
			continue
		} else if isodd == true {
			num, ok := <-even
			if ok == false {
				break
			}
			collector <- num
			iseven = true
			isodd = false
			continue
		}
	}
	close(collector)
}

func process_all(collector <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range collector {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup

	even := make(chan int)
	odd := make(chan int)
	collector := make(chan int, 3)

	wg.Add(1)
	go num_generator(odd, even, &wg)

	wg.Add(1)
	go proc_ev_odd(even, odd, collector, &wg)

	wg.Add(1)
	go process_all(collector, &wg)

	wg.Wait()
}
