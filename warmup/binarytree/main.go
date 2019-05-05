package main

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

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	println("node value: ", n.value)
	if value < n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
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

func main() {
	// https://en.wikipedia.org/wiki/Binary_search_tree
	// tree 8 3 10 1 6 14 4 7 13
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

	println(t.node.exists(30))
	println(t.node.exists(13))
	println(t.node.exists(1))
}
