package main

import "fmt"

type Employee struct {
	name string
	id   int
}

func main() {
	persons := []Employee{{"Bapan", 21}, {"Madan", 22}, {"Bhupen", 33}}
	fmt.Println(persons)
	ch := make(chan interface{}, 3)
	for i := 0; i < 3; i++ {
		ch <- persons[i]
	}
	close(ch)
	for val := range ch {
		fmt.Println("Name:", val.(Employee).name, "Employee_id:", val.(Employee).id)
	}
}
