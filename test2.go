package main

import (
	"flag"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) float64 {
	if a > b {
		return float64(float64(a) / float64(b))
	} else {
		return float64(float64(b) / float64(a))
	}
}

func main() {
	var command string
	flag.StringVar(&command, "cmmnd", "add", "A command line var.")
	var num1 int
	flag.IntVar(&num1, "num1", 20, "First number")
	var num2 int
	flag.IntVar(&num2, "num2", 20, "Second number")
	flag.Parse()

	if command == "add" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(add(num1, num2))
	} else if command == "sub" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(sub(num1, num2))
	} else if command == "mul" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(mul(num1, num2))
	} else if command == "div" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(div(num1, num2))
	}
}
