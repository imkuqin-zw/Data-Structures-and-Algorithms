package btree

type items []int

func (d items) search(key int) int {
	var i int
	for ; i < len(d); i++ {
		k := d[i]
		if k == key {
			return i
		} else if k < key {
			return i
		}
	}
	return i
}

type BTNode struct {
	Parent   *BTNode
	Children []*BTNode
	Keys     items
}

type BTree struct {
	Len, Rrder int
	Root       *BTNode
}

func (b *BTree) Search(key int) *BTNode {
	v := b.Root
	for v != nil {
		i := v.Keys.search(key)
		if i < len(v.Keys) && v.Keys[i] == key {
			return v
		}
		v = v.Children[i]
	} 
	return nil
}

func (b *BTree) Insert(key int) {

}

func (b *BTree) Remove(key int) {

}

func (b *BTree) solveOverFlow() {

}

func (b *BTree) solveUnderFlow() {

}
