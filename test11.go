package main

import "fmt"

func sum_num(nums ...int) int {
	sum := 0
	for val := range nums {
		sum += val
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum_num(nums...))
}
