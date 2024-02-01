package main

import "fmt"

func partition(arr *[10]int, start int, end int) int {
	pivot := (*arr)[end]
	wall := start - 1
	for i := start; i < end; i++ {
		if (*arr)[i] < pivot {
			wall += 1
			(*arr)[wall], (*arr)[i] = (*arr)[i], (*arr)[wall]
		}
	}
	wall += 1
	(*arr)[wall], (*arr)[end] = (*arr)[end], (*arr)[wall]
	return wall
}

func Qsort(arr *[10]int, start int, end int) {
	if start <= end {
		index := partition(arr, start, end)
		Qsort(arr, start, index-1)
		Qsort(arr, index+1, end)
	}
}

func main() {
	var arr [10]int
	arr = [10]int{12, 10, 3, 4, 13, 14, 9, 8, 2, 5}
	fmt.Println("Performing quick sort of the given array")
	Qsort(&arr, 0, 9)
	fmt.Println(arr)
}
