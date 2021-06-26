package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	node *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.Insert(value)
	}
	return t
}

func (n *Node) Insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.Insert(value)
		}
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	printNode(n.left)
	println(n.value)
	printNode(n.right)
}

func main() {
	t := &Tree{}
	t.Insert(10)
	t.Insert(3)
	t.Insert(7)
	t.Insert(17)
	printNode(t.node)
}
