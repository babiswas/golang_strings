package main

import (
	"container/list"
	"fmt"
)

func main() {
	lst := list.New()

	fmt.Println("Inserting value in linked list:")

	lst.PushBack(10)
	lst.PushBack(11)
	lst.PushBack(12)
	lst.PushBack(13)
	lst.PushBack(14)

	fmt.Println("Displaying numbers in linked list:")

	for ptr := lst.Front(); ptr != nil; ptr = ptr.Next() {
		fmt.Println(ptr.Value)
	}

	fmt.Println("Insert at the begining of the list:")

	lst.PushFront(8)
	lst.PushFront(7)
	lst.PushFront(6)

	fmt.Println("Displaying after entering the number at front:")

	for ptr := lst.Front(); ptr != nil; ptr = ptr.Next() {
		fmt.Println(ptr.Value)
	}

	fmt.Println("Displaying the length of the list")

	fmt.Println(lst.Len())

	fmt.Println()
}
