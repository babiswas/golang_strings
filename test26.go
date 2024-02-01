package main

import "fmt"

func bubble_sort(arr *[5]int) {
	for i := 4; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	var arr [5]int
	arr = [5]int{3, 4, 1, 7, 8}
	bubble_sort(&arr)
	fmt.Println(arr)
}
