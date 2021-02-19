package tree

type RBTree struct {
	Root   *RBNode
	Number int
}

type RBNode struct {
	Left  *RBNode
	Right *RBNode
	Color bool
	Key   int
	Value interface{}
}

func CreateRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) rotateLeft(node *RBNode) *RBNode {
	n := node.Right
	node.Right = n.Left
	n.Left = node
	node.Color = true
	return n
}

func (t *RBTree) rotateRight(node *RBNode) *RBNode {
	n := node.Left
	node.Left = n.Right
	n.Right = node
	n.Color = node.Color
	node.Color = true
	return n
}

func (t *RBTree) IsRed(node *RBNode) bool {
	if node == nil {
		return false
	}
	return node.Color
}

func (t *RBTree) flipColors(node *RBNode) {
	node.Left.Color = false
	node.Right.Color = false
	node.Color = true
}

func (t *RBTree) Put(key int, val interface{}) {
	t.Root = t.put(t.Root, key, val)
	t.Root.Color = false
}

func (t *RBTree) put(node *RBNode, key int, val interface{}) *RBNode {
	if node == nil {
		return &RBNode{Color: true, Key: key, Value: val}
	}
	if key < node.Key {
		node.Left = t.put(node.Left, key, val)
	} else if key > node.Key {
		node.Right = t.put(node.Right, key, val)
	} else {
		node.Value = val
	}
	if t.IsRed(node.Right) && !t.IsRed(node.Left) {
		node = t.rotateLeft(node)
	}
	if t.IsRed(node.Left) && t.IsRed(node.Left.Left) {
		node = t.rotateRight(node)
	}

	if t.IsRed(node.Left) && t.IsRed(node.Right) {
		t.flipColors(node)
	}

	return node
}

func (t *RBTree) Get(key int) interface{} {
	node := t.Root
	for node != nil {
		if node.Key == key {
			return node.Value
		} else if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
