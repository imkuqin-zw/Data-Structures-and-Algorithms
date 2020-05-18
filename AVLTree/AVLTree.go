package AVLTree

import "github.com/imkuqin-zw/Data-Structures-and-Algorithms/utils"

type AVLTree struct {
	Height int
	Val    int
	L, R   *AVLTree
}

func updateHeight(t *AVLTree) {
	if t.L == nil && t.R == nil {
		t.Height = 1
	} else if t.L == nil {
		t.Height = t.R.Height + 1
	} else if t.R == nil {
		t.Height = t.L.Height + 1
	} else {
		t.Height = utils.Max(t.L.Height, t.R.Height) + 1
	}
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

func LHAdjust(t *AVLTree) {
	if height(t.L)-height(t.R) >= 2 {
		if height(t.L.L) > height(t.L.R) {
			t = llRotation(t)
		} else {
			t = lrRotation(t)
		}
	}
}

func RHAdjust(t *AVLTree) {
	if height(t.L)-height(t.R) <= -2 {
		if height(t.R.R) > height(t.R.L) {
			t = rrRotation(t)
		} else {
			t = rlRotation(t)
		}
	}
}

func adjust(t *AVLTree) {
	LHAdjust(t)
	RHAdjust(t)
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
	if t.Val == val {
		return t
	}
	if val < t.Val {
		t.L = Insert(t.L, val)
		LHAdjust(t)
	} else {
		t.R = Insert(t.R, val)
		RHAdjust(t)
	}
	updateHeight(t)
	return t
}

func Delete(t *AVLTree, val int) *AVLTree {
	if t == nil {
		return nil
	}
	if t.Val == val {
		if t.R == nil {
			t = t.L
		} else {
			tmp := t.R
			for tmp.L != nil {
				tmp = tmp.L
			}
			t.Val = tmp.Val
			t.R = Delete(t.R, t.Val)
			updateHeight(t)
		}
		return t
	}
	if val < t.Val {
		t.L = Delete(t.L, val)
	} else {
		t.R = Delete(t.R, val)
	}
	if t.L != nil {
		adjust(t.L)
	}
	if t.R != nil {
		adjust(t.R)
	}
	updateHeight(t)
	return t
}
