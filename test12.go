package main

import "fmt"

func main() {
	test := "hello"
	for _, val := range test {
		fmt.Printf("%c\n", val)
	}
}
