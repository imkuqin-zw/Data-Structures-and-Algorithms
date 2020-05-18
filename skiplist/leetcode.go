package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxLevel = 32
const Probability = 0.25

type SkipNode struct {
	num  int
	next []*SkipNode
	prev []*SkipNode
}

type Skiplist struct {
	head  *SkipNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head: &SkipNode{
			num:  -1,
			next: make([]*SkipNode, MaxLevel),
		},
		level: 1,
	}
}

func (this *Skiplist) randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	level = 1
	for rand.Float32() < Probability && level < MaxLevel {
		level++
	}
	return MaxLevel
}

func (this *Skiplist) Search(target int) bool {
	p := this.head
	for i := this.level - 1; i >= 0; i-- {
		for p.next[i] != nil {
			if p.next[i].num == target {
				return true
			} else if p.next[i].num < target {
				p = p.next[i]
			} else {
				break
			}
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	p := this.head
	update := make([]*SkipNode, this.level)
	for i := this.level - 1; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].num < num {
			p = p.next[i]
		}
		update[i] = p
	}

	if p.next[0] != nil && p.next[0].num == num {
		return
	}

	level := this.randLevel()
	if level > this.level {
		this.level++
		level = this.level
		update = append(update, this.head)
	}
	node := &SkipNode{
		num:  num,
		next: make([]*SkipNode, level),
		prev: make([]*SkipNode, level),
	}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		node.prev[i] = update[i]
		update[i].next[i] = node
	}
}

func (this *Skiplist) Erase(num int) bool {
	p := this.head
	for i := this.level - 1; i >= 0; i-- {
		for p.next[i] != nil {
			if p.next[i].num == num {
				for {
					tmp := p.next[i]
					p.next[i] = tmp.next[i]
					if tmp.next[i] != nil {
						tmp.next[i].prev[i] = p
					}
					tmp.next[i] = nil
					tmp.prev[i] = nil
					i--
					if i < 0 {
						for this.level > 1 && this.head.next[this.level-1] == nil {
							this.level--
						}
						return true
					}
					p = tmp.prev[i]
				}
			} else if p.next[i].num < num {
				p = p.next[i]
			} else {
				break
			}
		}
	}

	return false
}

func (this *Skiplist) println() {
	for i := this.level; i >= 0; i-- {
		p := this.head
		fmt.Print(p.num, " ")
		for p.next[i] != nil {
			fmt.Print(p.next[i].num, " ")
			p = p.next[i]
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide
}

func min(i, j, k int) int {
	m := i
	if m > j {
		m = j
	}
	if m > k {
		m = k
	}
	return m
}

func climbStairs(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]) + 1
	}
	return dp[n]
}
