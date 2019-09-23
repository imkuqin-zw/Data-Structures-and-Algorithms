package main

import (
	"testing"
)

func TestBTNode_insert(t *testing.T) {
	var bt *BTNode
	for i := 1; i < 10; i++ {
		bt = bt.insert(i)
	}
	bt = bt.insert(0)
	bt = bt.remove(1)
}

func TestAVL_insert(t *testing.T) {
	var avl *AVL
	for i := 1; i < 10; i++ {
		avl.insert(i)
	}
}
