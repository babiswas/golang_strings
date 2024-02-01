package main

import "fmt"

func insertion_sort(arr *[10]int, start int) {

	for i := start + 1; i < 10; i++ {
		val := (*arr)[i]
		hole := i
		for ((*arr)[hole-1] > val) && (hole > 0) {
			(*arr)[hole] = (*arr)[hole-1]
			hole = hole - 1
			if hole == 0 {
				break
			}
		}
		(*arr)[hole] = val
	}
}

func main() {
	var arr [10]int
	arr = [10]int{11, 5, 1, 16, 2, 6, 7, 0, 21, -2}
	insertion_sort(&arr, 0)
	fmt.Println(arr)
}
