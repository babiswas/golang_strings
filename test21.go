package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(200))
	fmt.Println((r1.Float64() * 5) + 5)
}
