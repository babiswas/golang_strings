package main

import (
	"flag"
	"fmt"
)

func add_numbers(a, b int) int {
	return a + b
}

func substract_numbers(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func multiply_numbers(a, b int) int {
	return a * b
}

func divide_numbers(a, b int) float64 {
	if a > b {
		return float64(float64(a) / float64(b))
	} else {
		return float64(float64(b) / float64(a))
	}
}

func main() {

	fmt.Println("Calculator:")
	fmt.Println("===========================================================")
	command_str := flag.String("op_cmd", "add", "Arithmetic op[eration.")
	num1 := flag.Int("num1", 20, "First number")
	num2 := flag.Int("num2", 21, "Second number")

	flag.Parse()

	command := *command_str
	first_number := *num1
	second_number := *num2

	if command == "add" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(add_numbers(first_number, second_number))
	} else if command == "sub" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(substract_numbers(first_number, second_number))
	} else if command == "mul" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(multiply_numbers(first_number, second_number))
	} else if command == "div" {
		fmt.Println("The sum of the numbers is:")
		fmt.Println(divide_numbers(first_number, second_number))
	} else {
		fmt.Println("Command not known.")
	}
	fmt.Println("===========================================================")
}
