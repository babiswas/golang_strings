package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 9, 1, 8, 11}
	sort.Ints(nums)
	fmt.Println(nums)
	names := []string{"hello", "tello", "apple", "cat"}
	sort.Strings(names)
	fmt.Println(names)
	mynums := []float64{6.4, 2.2, 1.4, 7.2, 6.4}
	sort.Float64s(mynums)
	fmt.Println(mynums)
}
