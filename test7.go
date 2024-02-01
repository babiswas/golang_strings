package main

import (
	"context"
	"fmt"
	"time"
)

func do_work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Done")
			return
		default:
			fmt.Println("Working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go do_work(ctx)
	select {
	case <-ctx.Done():
	case <-time.After(3 * time.Second):
		cancel()
	}
}
