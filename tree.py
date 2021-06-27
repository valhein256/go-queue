class Node():
    def __init__ (self, value: int):
        self.value = value
        self.left = None
        self.right = None

    def insert(self, value: int):
        if self.value > value:
            if self.left:
                self.left.insert(value)
            else:
                self.left = Node(value)
        else:
            if self.right:
                self.right.insert(value)
            else:
                self.right = Node(value)

class Tree():
    node = None

    def insert(self, value: int):
        if self.node == None:
            self.node = Node(value)
        else:
            self.node.insert(value)

def printNode(node: Node):
    if node == None:
        return

    printNode(node.left)
    print(node.value)
    printNode(node.right)

t = Tree()
t.insert(10)
t.insert(3)
t.insert(4)
t.insert(23)

printNode(t.node)
