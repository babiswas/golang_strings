package main

import (
	"flag"
	"fmt"
)

func is_palindrome(str1 string) bool {
	index1 := 0
	index2 := len(str1) - 1
	for index1 <= index2 {
		if str1[index1] == str1[index2] {
			index1 += 1
			index2 -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	first_string := flag.String("string1", "", "a string")
	flag.Parse()
	string1 := *first_string
	if is_palindrome(string1) {
		message := fmt.Sprintf("%s string is palindrome", string1)
		fmt.Println(message)
	} else {
		message := fmt.Sprintf("%s is not palindrome.", string1)
		fmt.Println(message)
	}
}
