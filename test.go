package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	ch := greet(quit)
	num := rand.Intn(10000000)
	for i := 0; i <= num; i++ {
		fmt.Println("Message received ", <-ch)
	}
	quit <- true
	fmt.Println("Quit the routine.")
}

func greet(quit <-chan bool) <-chan string {
	ch1 := make(chan string)
	go func() {
		for {
			select {
			case ch1 <- fmt.Sprintf("Hello %s", "bapan"):
				fmt.Println("Passed a message.")
			case <-quit:
				fmt.Println("Quitting the routine")
				return
			}
		}
	}()
	return ch1
}
