package main

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
	t.insert(2)
	t.insert(1)
	t.insert(3)
	t.insert(7)
	t.insert(23)
	printTree(t.node)

	println(t.node.search(10))
	println(t.node.search(12))
}
