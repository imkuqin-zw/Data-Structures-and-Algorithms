package stack

import (
	"container/list"
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/stack/arrayStack"
	"strconv"
)

func MvBottomToTop(s Stack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		MvBottomToTop(s)
		top2 := s.Pop()
		s.Push(top1)
		s.Push(top2)
	} else {
		s.Push(top1)
	}
}

func ReverseStack(s Stack) {
	if s.IsEmpty() {
		return
	}
	MvBottomToTop(s)
	top := s.Pop()
	ReverseStack(s)
	s.Push(top)
}

func MvBottomToTopSort(s Stack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		MvBottomToTopSort(s)
		top2 := s.Top()
		if top2.(int) < top1.(int) {
			s.Pop()
			s.Push(top1)
			s.Push(top2)
			return
		}
	}
	s.Push(top1)
}

func SortStack(s Stack) {
	if s.IsEmpty() {
		return
	}
	MvBottomToTopSort(s)
	top := s.Pop()
	SortStack(s)
	s.Push(top)
}

func IsPopSerial(in, out []int) bool {
	inLen, outLen := len(in), len(out)
	if inLen != outLen {
		return false
	}
	s := arrayStack.NewStack()
	inIndex, outIndex := 0, 0
	for inIndex < inLen {
		s.Push(in[inIndex])
		inIndex++
		for !s.IsEmpty() && s.Top().(int) == out[outIndex] {
			s.Pop()
			outIndex++
		}
	}
	if !s.IsEmpty() || outIndex != outLen {
		return false
	}
	return true
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	op := map[byte]byte{
		'(': ')',
		')': '(',
		'{': '}',
		'}': '{',
		'[': ']',
		']': '[',
	}
	stack := list.New()
	for i := 0; i < len(s); i++ {
		if stack.Len() == 0 {
			stack.PushBack(s[i])
		} else {
			e := stack.Back()
			if op[e.Value.(byte)] == s[i] {
				stack.Remove(e)
			} else {
				stack.PushBack(s[i])
			}
		}
	}
	return stack.Len() == 0
}

func dailyTemperatures(T []int) []int {
	stack := list.New()
	type Node struct {
		t   int
		idx int
	}
	rsp := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		node := &Node{t: T[i], idx: i}
		for stack.Len() > 0 {
			e := stack.Back()
			prev := e.Value.(*Node)
			if prev.t >= node.t {
				break
			}
			rsp[prev.idx] = i - prev.idx
			stack.Remove(e)
		}
		stack.PushBack(node)
	}
	return rsp
}

func evalRPN(tokens []string) int {
	op := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stack := list.New()
	for i := 0; i < len(tokens); i++ {
		var num int
		if !op[tokens[i]] {
			num, _ = strconv.Atoi(tokens[i])
		} else {
			n2E := stack.Back()
			stack.Remove(n2E)
			n1E := stack.Back()
			stack.Remove(n1E)
			n1, n2 := n1E.Value.(int), n2E.Value.(int)
			switch tokens[i] {
			case "+":
				num = n1 + n2
			case "-":
				num = n1 - n2
			case "*":
				num = n1 * n2
			case "/":
				num = n1 / n2
			}
		}
		stack.PushBack(num)
	}
	return stack.Back().Value.(int)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	type node struct {
		i, j int
	}
	num := 0
	stack := list.New()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0'
			stack.PushBack(&node{i, j})
			for stack.Len() > 0 {
				e := stack.Back()
				stack.Remove(e)
				n := e.Value.(*node)
				if n.j+1 < len(grid[i]) && grid[n.i][n.j+1] == '1' {
					grid[n.i][n.j+1] = '0'
					stack.PushBack(&node{n.i, n.j + 1})
				}
				if n.j-1 >= 0 && grid[n.i][n.j-1] == '1' {
					grid[n.i][n.j-1] = '0'
					stack.PushBack(&node{n.i, n.j - 1})
				}
				if n.i+1 < len(grid) && grid[n.i+1][n.j] == '1' {
					grid[n.i+1][n.j] = '0'
					stack.PushBack(&node{n.i + 1, n.j})
				}
				if n.i-1 >= 0 && grid[n.i-1][n.j] == '1' {
					grid[n.i-1][n.j] = '0'
					stack.PushBack(&node{n.i - 1, n.j})
				}
			}
			num++
		}
	}
	return num
}

type Node struct {
	Val       int
	Neighbors []*Node
}

var hashMap = make(map[*Node]*Node)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	result := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	hashMap[node] = result
	for i := range node.Neighbors {
		if n, ok := hashMap[node.Neighbors[i]]; !ok {
			result.Neighbors[i] = cloneGraph(node.Neighbors[i])
		} else {
			result.Neighbors[i] = n
		}
	}
	return result
}

func subarraySum(nums []int, k int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		curSum := 0
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			if curSum == k {
				n++
			}
		}
	}
	return n
}

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back()
		node := e.Value.(*TreeNode)
		if node.Left != nil {
			stack.PushBack(node.Left)
			node.Left = nil
			continue
		}
		result = append(result, node.Val)
		stack.Remove(e)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	return result
}

func pivotIndex(nums []int) int {
	right := 0
	for i := 0; i < len(nums); i++ {
		right += nums[i]
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}

func dominantIndex(nums []int) int {
	first := []int{-1, -1}
	second := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] > first[0] {
			second[0], second[1] = first[0], first[1]
			first[0], first[1] = nums[i], i
		} else if nums[i] > second[0] {
			second[0], second[1] = nums[i], i
		}
	}
	if second[0]*2 <= first[0] {
		return first[1]
	}
	return -1
}

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] -= 10
		} else {
			return digits
		}
	}
	if i == -1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
