package main

import (
	"context"
	"fmt"
	"time"
)

func mygreet(ctx context.Context) <-chan string {
	ch1 := make(chan string)
	go func() {
		for {
			select {
			case ch1 <- fmt.Sprintf("Hello %s", "bapan"):
				fmt.Println("Message Passed.")
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					fmt.Println(err)
				}
				fmt.Println("Quitting the routine.")
				return
			}
		}
	}()
	return ch1
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch1 := mygreet(ctx)
	num := 100
	for i := 0; i < num; i++ {
		fmt.Println("Message received ", <-ch1)
	}
	cancel()
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Quiting the process.")
}
