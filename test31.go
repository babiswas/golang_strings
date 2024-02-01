package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

func (n *Node) add_node(num int) *Node {
	return &Node{num, nil}
}

type List struct {
	list_node *Node
}

func (l *List) add_last(num int) {
	node := Node{}
	if l.list_node == nil {
		l.list_node = node.add_node(num)
	} else {
		temp := l.list_node
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node.add_node(num)
	}
}

func (l *List) add_front(num int) {
	node := Node{}
	if l.list_node == nil {
		l.list_node = node.add_node(num)
	} else {
		temp := l.list_node
		l.list_node = node.add_node(num)
		l.list_node.next = temp
	}
}

func (l *List) display() {
	temp := l.list_node
	for temp != nil {
		fmt.Println(temp.num)
		temp = temp.next
	}
}

func main() {
	l := List{}
	fmt.Println("Displaying after adding elements in back:")
	l.add_last(2)
	l.add_last(3)
	l.add_last(5)
	l.add_last(6)
	l.display()
	fmt.Println("Displaying after adding elements in front")
	l.add_front(10)
	l.add_front(11)
	l.add_front(12)
	l.display()
}
