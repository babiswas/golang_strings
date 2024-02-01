package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5}
	n := make([]int, len(m))
	copy(n, m)
	fmt.Println(n)
}
