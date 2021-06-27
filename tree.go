package main

import "fmt"

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

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
}

func (n *Node) search(value int) int {
	if n == nil {
		return -1
	}

	if value == n.value {
		return n.value
	}

	if value < n.value {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}

func printTree(n *Node, sort_type string) {
	if n == nil {
		return
	}
	switch sort_type {
	case "ascending":
		printTree(n.left, sort_type)
		println(n.value)
		printTree(n.right, sort_type)
		break
	case "descending":
		printTree(n.right, sort_type)
		println(n.value)
		printTree(n.left, sort_type)
		break
	default:
		println(n.value)
		printTree(n.left, sort_type)
		printTree(n.right, sort_type)
		break
	}
}

func main() {
	t := &Tree{}
	t.insert(10)
	t.insert(3)
	t.insert(2)
	t.insert(1)
	t.insert(3)
	t.insert(7)
	t.insert(23)
	println("ascending")
	printTree(t.node, "ascending")
	println("descending")
	printTree(t.node, "descending")
	println("default")
	printTree(t.node, "")
	println()

	for i := 1; i <= 10; i++ {
		r := t.node.search(i)
		if r > 0 {
			fmt.Printf("%d is existed in Tree\n", i)
		} else {
			fmt.Printf("%d is not existed in Tree\n", i)
		}
	}
}
