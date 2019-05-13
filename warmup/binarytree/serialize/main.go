package main

import "fmt"

type Tree struct {
	node *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	println(n.value)
	printNode(n.left)
	printNode(n.right)
}

func serialize(n *Node, out *[]int) {
	if n == nil {
		return
	}
	*out = append(*out, n.value)
	serialize(n.left, out)
	serialize(n.right, out)
}

func main() {
	t := &Tree{}
	t.Insert(8).
		Insert(3).
		Insert(10).
		Insert(1).
		Insert(6).
		Insert(14).
		Insert(4).
		Insert(7).
		Insert(13)
	printNode(t.node)

	out := &[]int{}
	serialize(t.node, out)
	fmt.Printf("out = %+v\n", out)
}
