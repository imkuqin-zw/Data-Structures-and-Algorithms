package BTree

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/queue/listQueue"
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/listStack"
)

type BNode struct {
	Data                   interface{}
	Parent, LChild, RChild *BNode
	Height                 int
	Npl                    int
	Color                  int
}

func NewBNode(data interface{}) *BNode {
	return &BNode{Data: data}
}

func (b *BNode) visitAlongLeftBranch(out []interface{}, stack *listStack.Stack) {
	out = append(out, b.Data)
	cur := b.LChild
	for cur != nil {
		out = append(out, cur.Data)
		if cur.RChild != nil {
			stack.Push(cur.RChild)
		}
	}
}

func (b *BNode) PreOrderTraversal() []interface{} {
	out := make([]interface{}, 0)
	if b == nil {
		return out
	}
	nodeStack := listStack.NewStack()
	root := b
	for {
		root.visitAlongLeftBranch(out, nodeStack)
		if nodeStack.IsEmpty() {
			break
		}
		root = nodeStack.Top().(*BNode)
	}
	return out
}

func (b *BNode) goAlongLeftBranch(out []interface{}, stack *listStack.Stack) {
	cur := b
	for cur != nil {
		stack.Push(cur)
		cur = cur.LChild
	}
}

func (b *BNode) InOrderTraversal() []interface{} {
	out := make([]interface{}, 0)
	if b == nil {
		return out
	}
	nodeStack := listStack.NewStack()
	root := b
	for {
		root.goAlongLeftBranch(out, nodeStack)
		if nodeStack.IsEmpty() {
			break
		}
		root = nodeStack.Top().(*BNode)
		out = append(out, root.Data)
		root = root.RChild
	}
	return out
}

func (b *BNode) afterAlongLeftBranch(out []interface{}, stack *listStack.Stack) {
	cur := b
	for cur != nil {
		if cur.LChild == nil && cur.RChild == nil {
			out = append(out, cur.Data)
			break
		}
		if cur.RChild != nil {
			stack.Push(cur.RChild)
		}
		tmp := cur
		cur = cur.LChild
		tmp.LChild = nil
		stack.Push(tmp)
	}
}

func (b *BNode) PostOrderTraversal() []interface{} {
	out := make([]interface{}, 0)
	if b == nil {
		return out
	}
	nodeStack := listStack.NewStack()
	root := b
	for {
		root.afterAlongLeftBranch(out, nodeStack)
		if nodeStack.IsEmpty() {
			break
		}
		root = nodeStack.Top().(*BNode)
	}
	return out
}

func (b *BNode) LevelTraversal() []interface{} {
	out := make([]interface{}, 0)
	que := listQueue.NewQueue()
	que.Enqueue(b)
	for !que.Empty() {
		node := que.Dequeue().(*BNode)
		out = append(out, node.Data)
		if node.LChild != nil {
			que.Enqueue(node.LChild)
		}
		if node.RChild != nil {
			que.Enqueue(node.RChild)
		}
	}
	return out
}
