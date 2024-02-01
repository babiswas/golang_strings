package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func process_cntxt(ctx context.Context) {
	val := ctx.Value("hello")
	fmt.Println(val)
}

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
