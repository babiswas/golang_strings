package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(time.Second + 5)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second * 3):
		fmt.Println("Over Slept.")
	}
}
