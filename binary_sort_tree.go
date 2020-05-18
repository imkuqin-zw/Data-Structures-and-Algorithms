package main

// AVLTree is
type AVL struct {
	data   int
	height int
	l, r   *AVL
}

func NewAVL(data int) *AVL {
	return &AVL{data: data, height: 1}
}

func (avl *AVL) rrRotate() *AVL {
	newTop := avl.l
	avl.l = newTop.r
	newTop.r = avl
	avl.height = max(getHeight(avl.l), getHeight(avl.r)) + 1
	newTop.height = max(getHeight(avl.l), getHeight(avl.r)) + 1
	return newTop
}

func (avl *AVL) lrRotate() *AVL {
	avl.l = avl.l.llRotate()
	return avl.rrRotate()
}

func (avl *AVL) rlRotate() *AVL {
	avl.r = avl.r.rrRotate()
	return avl.llRotate()
}

func (avl *AVL) llRotate() *AVL {
	newTop := avl.r
	avl.r = newTop.l
	newTop.l = avl
	return newTop
}

func (avl *AVL) balance() *AVL {
	bf := getHeight(avl.l) - getHeight(avl.r)
	if bf == 2 {
		if getHeight(avl.l.l)-getHeight(avl.l.r) > 0 {
			return avl.llRotate()
		} else {
			return avl.lrRotate()
		}
	} else if bf == -2 {
		if getHeight(avl.r.l)-getHeight(avl.r.r) > 0 {
			return avl.rlRotate()
		} else {
			return avl.rrRotate()
		}
	}
	return avl
}

func (avl *AVL) insert(data int) *AVL {
	if avl == nil {
		return NewAVL(data)
	}
	var result *AVL
	if avl.data > data {
		avl.l = avl.l.insert(data)
		result = avl.balance()
	} else {
		avl.r = avl.r.insert(data)
		result = avl.balance()
	}
	result.height = max(getHeight(avl.l), getHeight(avl.r))
	return result
}

func getHeight(bn *AVL) int {
	if bn == nil {
		return 0
	}
	return bn.height
}

// BTNode is one node of the binary tree.
type BTNode struct {
	data   int
	l, r   *BTNode
	height int
}

func getBTHeight(bn *BTNode) int {
	if bn == nil {
		return 0
	}
	return bn.height
}

func (bn *BTNode) contain(data int) bool {
	if bn == nil {
		return false
	}
	if bn.data < data {
		return bn.r.contain(data)
	} else if bn.data > data {
		return bn.l.contain(data)
	}
	return true
}

func (bn *BTNode) insert(data int) *BTNode {
	if bn == nil {
		return &BTNode{data: data, height: 1}
	}
	if data < bn.data {
		bn.l = bn.l.insert(data)
	} else {
		bn.r = bn.r.insert(data)
	}
	bn.height = max(getBTHeight(bn.l), getBTHeight(bn.r)) + 1
	return bn
}

func (bn *BTNode) remove(data int) *BTNode {
	if bn == nil {
		return bn
	}
	if data == bn.data {
		if bn.l == nil {
			return nil
		} else if bn.r == nil {
			return bn.l
		} else {
			bn.data = bn.r.findMin().(int)
			bn.r = bn.r.remove(bn.data)
		}
		return bn
	} else if data < bn.data {
		bn.l = bn.l.remove(data)
	} else {
		bn.r = bn.r.remove(data)
	}
	return bn
}

func (bn *BTNode) findMin() interface{} {
	if bn == nil {
		return nil
	}
	if bn.l == nil {
		return bn.data
	}
	return bn.l.findMin()
}

func (bn *BTNode) findMax() interface{} {
	if bn == nil {
		return nil
	}
	cur := bn
	for cur.r != nil {
		cur = cur.r
	}
	return cur.data
}
