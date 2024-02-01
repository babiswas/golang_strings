package main

import "fmt"

func selection_sort(arr *[5]int) {
	//Selection sort of an array in golang
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
}

func main() {
	var arr [5]int
	arr = [5]int{3, 6, 7, 1, 11}
	selection_sort(&arr)
	fmt.Println(arr)
}
