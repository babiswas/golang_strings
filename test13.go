package main

import "fmt"

func reverse(s string) string {
	val := []rune(s)
	start := 0
	end := len(s) - 1
	for start <= end {
		temp := val[start]
		val[start] = val[end]
		val[end] = temp
		start += 1
		end -= 1
	}
	return string(val)
}

func main() {
	temp := "hello world"
	fmt.Println(temp)
	fmt.Println("Reversing the string:")
	str1 := reverse(temp)
	fmt.Println(str1)
}
