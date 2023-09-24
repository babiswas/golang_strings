package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	l.PushBack(3)
	l.PushFront(10)
	l.InsertAfter(11, e1)
	l.InsertBefore(34, e2)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
