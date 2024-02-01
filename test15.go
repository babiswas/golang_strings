package main

import "fmt"

func remove_index(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	test := make([]int, 0)
	index := 0
	for index != 5 {
		test = append(test, 10+index)
		index += 1
	}
	fmt.Println(test)
	test = remove_index(test, 2)
	fmt.Println(test)
}
