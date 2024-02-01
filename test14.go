package main

import "fmt"

type stack []int

func (s *stack) push(num int) {
	if *s == nil {
		*s = make([]int, 0)
	}
	*s = append(*s, num)
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[0:index]
	return item
}

func (s stack) display_stack() {
	fmt.Println(s)
}

func main() {
	var s stack
	s.push(1)
	s.push(2)
	s.push(3)
	s.display_stack()
	val := s.pop()
	for val != -1 {
		fmt.Println(val)
		val = s.pop()
	}
}
