package main

type Node struct {
	value int
	left  *Node
	right *Node
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

func printTree(n *Node) {
	if n == nil {
		return
	}
	printTree(n.left)
	println(n.value)
	printTree(n.right)
}

func main() {
	t := &Tree{}
	t.insert(10)
	t.insert(3)
	t.insert(7)
	t.insert(23)
	printTree(t.node)
}
