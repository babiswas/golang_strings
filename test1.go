package main

import (
	"fmt"
	"log"
)

func main() {
	m := make(map[int]bool)

	log.Println("Adding values in set.")

	add := func(num int) {
		_, ok := m[num]
		if ok == false {
			m[num] = true
		}
	}

	log.Println("Removing values in set.")

	remove := func(num int) {
		_, ok := m[num]
		if ok == true {
			delete(m, num)
		}
	}

	log.Println("Does the number exist.")

	is_exist := func(num int) bool {
		_, ok := m[num]
		if ok == true {
			return true
		} else {
			return false
		}
	}

	m[1] = true
	m[2] = true
	m[3] = true
	m[4] = true

	fmt.Println("============================================")
	fmt.Println("Map operations:")

	fmt.Println("===========CHECK IF EXIST====================")

	for _, val := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		message := fmt.Sprintf("Does the number %d exist", val)
		log.Println(message)
		fmt.Println("Does the number exist ", is_exist(val))
	}

	fmt.Println("================================================")

	fmt.Println("===============INSERT NUMBER====================")

	for _, val := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		message := fmt.Sprintf("Insert the number %d", val)
		log.Println(message)
		add(val)
	}

	fmt.Println("================================================")

	fmt.Println("=================REMOVE NUMBER===================")

	for _, val := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		message := fmt.Sprintf("Remove the number %d", val)
		log.Println(message)
		remove(val)
	}

	fmt.Println("==================================================")

}
