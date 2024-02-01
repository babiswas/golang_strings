package main

import (
	"fmt"
	"sync"
)

func number_generator(odd chan<- int, even chan<- int, wg *sync.WaitGroup) {
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

func process(even <-chan int, odd <-chan int, wg *sync.WaitGroup) {
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
			fmt.Println(num)
			isodd = true
			iseven = false
			continue
		} else if isodd == true {
			num, ok := <-even
			if ok == false {
				break
			}
			fmt.Println(num)
			iseven = true
			isodd = false
			continue
		}
	}
}

func main() {
	var wg sync.WaitGroup

	even := make(chan int)
	odd := make(chan int)

	wg.Add(1)
	go number_generator(odd, even, &wg)

	wg.Add(1)
	go process(even, odd, &wg)

	wg.Wait()
}
