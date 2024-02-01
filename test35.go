package main

import "fmt"

type Leaf struct {
	val   int
	left  *Leaf
	right *Leaf
}

func (n *Leaf) create_leaf(value int) {
	n.val = value
	n.left = nil
	n.right = nil
}

type Tree struct {
	leaf *Leaf
}

func (leaf *Leaf) insert(value int) {
	newLeaf := Leaf{}
	newLeaf.create_leaf(value)
	if leaf.val > value {
		if leaf.right == nil {
			leaf.right = &newLeaf
		} else {
			leaf.right.insert(value)
		}
	} else if leaf.val < value {
		if leaf.left == nil {
			leaf.left = &newLeaf
		} else {
			leaf.left.insert(value)
		}
	} else {
		fmt.Println("Duplicate key:", value)
	}
}

func (t *Tree) insert_leaf(value int) {
	if t.leaf == nil {
		newLeaf := Leaf{}
		newLeaf.create_leaf(value)
		t.leaf = &newLeaf
	} else {
		t.leaf.insert(value)
	}
}
func (leaf *Leaf) inorder() {
	if leaf.left != nil {
		leaf.left.inorder()
	}
	fmt.Println(leaf.val)
	if leaf.right == nil {
		leaf.right.inorder()
	}
}
func (t *Tree) traverse_tree() {
	if t.leaf == nil {
		return
	} else {
		t.leaf.inorder()
	}
}
func main(){
	t:=Tree{}

	fmt.Println("Building the tree:")
	t.insert_leaf(32)
	t.insert_leaf(21)
	t.insert_leaf(44)
	t.insert_leaf(52)
	t.insert_leaf(40)
	t.insert_leaf(19)
	t.insert_leaf(27)

	fmt.Println("Tree traversal:")
	t.traverse_tree()
}

