package AVLTree

type AVLTree struct {
	Height int
	Val    int
	L, R   *AVLTree
}

func updateHeight(t *AVLTree) {

}

func llRotation(t *AVLTree) *AVLTree {
	tmp := t.L
	t.L = tmp.R
	tmp.R = t
	updateHeight(t)
	updateHeight(tmp)
	return tmp
}

func rrRotation(t *AVLTree) *AVLTree {
	tmp := t.R
	t.R = tmp.L
	tmp.L = t
	updateHeight(t)
	updateHeight(tmp)
	return tmp
}

func lrRotation(t *AVLTree) *AVLTree {
	t.L = rrRotation(t.L)
	return llRotation(t)
}

func rlRotation(t *AVLTree) *AVLTree {
	t.R = llRotation(t.R)
	return rrRotation(t)
}
func height(t *AVLTree) int {
	if t == nil {
		return 0
	}
	return t.Height
}

func adjust(t *AVLTree) *AVLTree {
	if height(t.L) - height(t.R) >= 2 {
		if height(t.L.L) > height(t.L.R) {
			t = llRotation(t)
		} else {
			t = lrRotation(t)
		}
	}
	if height(t.L) - height(t.R) <= -2 {
		if height(t.R)
	}
}

func Insert(t *AVLTree, val int) *AVLTree {
	if t == nil {
		t = &AVLTree{
			Height: 1,
			Val:    val,
			L:      nil,
			R:      nil,
		}
		return t
	}

}
