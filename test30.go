package main

import "fmt"

func rotate_arr(arr *[10]int, start int, end int, num int) {
	for start <= end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start = start + 1
		end = end - 1
	}
	start_index := 0
	end_index := num - 1

	for start_index <= end_index {
		(*arr)[start_index], (*arr)[end_index] = (*arr)[end_index], (*arr)[start_index]
		start_index += 1
		end_index -= 1
	}

	start_index = num
	end_index = len(*arr) - 1

	for start_index <= end_index {
		(*arr)[start_index], (*arr)[end_index] = (*arr)[end_index], (*arr)[start_index]
		start_index += 1
		end_index -= 1
	}

}

func main() {
	brr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(brr)
	rotate_arr(&brr, 0, len(brr)-1, 3)
	fmt.Println(brr)
}
