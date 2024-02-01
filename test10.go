package main

import (
	"context"
	"fmt"
	"time"
)

func myfunc(ctx context.Context) <-chan string {
	ch1 := make(chan string)
	go func() {
		for {
			select {
			case ch1 <- fmt.Sprintf("Hello Bapan"):
				fmt.Println("Message passed")
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					fmt.Println(err)
				}
				fmt.Println("Execution Completed")
				return
			}

		}
	}()
	return ch1
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch1 := myfunc(ctx)
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch1)
	}
	cancel()
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Quitting the process.")
}
