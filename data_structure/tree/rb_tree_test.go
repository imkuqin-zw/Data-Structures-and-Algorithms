package tree

import (
	"fmt"
	"testing"
)

func TestRBTree(t *testing.T) {
	rbTree := CreateRBTree()
	rbTree.Put(1, "zhang")
	rbTree.Put(2, "li")
	rbTree.Put(3, "wang")

	fmt.Println(rbTree.Get(1))
	fmt.Println(rbTree.Get(2))
	fmt.Println(rbTree.Get(3))
}
