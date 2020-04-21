package binary_tree

import "github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/listStack"

func FindMaxSubTree(root *BNode) (int, *BNode) {
	if root == nil {
		return 0, root
	}
	type loopFunc func(*BNode, **BNode) int
	var fn loopFunc
	var maxSum int
	fn = func(node *BNode, maxRoot **BNode) int {
		if node == nil {
			return 0
		}
		lMax := fn(node.LChild, maxRoot)
		rMax := fn(node.RChild, maxRoot)
		sum := lMax + rMax + node.Data.(int)
		if sum > maxSum {
			maxSum = sum
			*maxRoot = node
		}
		return sum
	}
	var maxRoot *BNode
	fn(root, &maxRoot)
	return maxSum, maxRoot
}

func IsEqual(root1, root2 *BNode, cmp func(va1, va2 interface{}) int) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		if cmp(root1.Data, root2.Data) == 0 {
			return IsEqual(root1.LChild, root2.LChild, cmp) && IsEqual(root1.RChild, root2.RChild, cmp)
		}
	}
	return false
}

func BTreeToDLink(b *BNode) *BNode {
	if b == nil {
		return nil
	}
	var head, end *BNode
	s := listStack.NewStack()
	cur := b
	for {
		if cur != nil {
			s.Push(cur)
			cur = cur.LChild
		} else {
			if s.IsEmpty() {
				break
			}
			node := s.Pop().(*BNode)
			node.LChild = end
			if end == nil {
				head = node
			} else {
				end.RChild = node
			}
			end = node
			cur = node.RChild
		}
	}
	return head
}

func FindParentNode(root, node1, node2 *BNode) *BNode {
	if root == nil || node1.Data == root.Data || node2.Data == root.Data {
		return root
	}
	l := FindParentNode(root.LChild, node1, node2)
	r := FindParentNode(root.RChild, node1, node2)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}
