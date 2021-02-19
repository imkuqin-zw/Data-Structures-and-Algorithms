package tree

type UFTree struct {
	EAP   []int
	ES    []int
	count int
}

func NewUFTree(num int) *UFTree {
	ufTree := &UFTree{
		EAP:   make([]int, num),
		ES:    make([]int, num),
		count: num,
	}
	for i := 0; i < num; i++ {
		ufTree.ES[i] = 1
	}
	for i := 0; i < num; i++ {
		ufTree.EAP[i] = i
	}
	return ufTree
}

func (t *UFTree) Find(p int) int {
	for p != t.EAP[p] {
		p = t.EAP[p]
	}
	return p
}

func (t *UFTree) Connected(p, q int) bool {
	return t.Find(p) == t.Find(q)
}

func (t *UFTree) Union(p, q int) {
	pRoot := t.Find(p)
	qRoot := t.Find(q)

	if pRoot == qRoot {
		return
	}

	if t.ES[pRoot] > t.ES[qRoot] {
		t.EAP[qRoot] = t.EAP[pRoot]
		t.ES[qRoot] += t.ES[pRoot]
	} else {
		t.EAP[pRoot] = t.EAP[qRoot]
		t.ES[pRoot] = t.ES[qRoot]
	}

	t.count--
}

func (t *UFTree) Count() int {
	return t.count
}
