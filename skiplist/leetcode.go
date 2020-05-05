package skiplist

import (
	"math/rand"
	"time"
)

const MaxLevel = 32
const Probability = 0.25

type SkipNode struct {
	num  int
	prev []*SkipNode
	next []*SkipNode
}

type Skiplist struct {
	head  *SkipNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head: &SkipNode{
			num:  -1,
			prev: make([]*SkipNode, MaxLevel),
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
	return
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
	level := this.randLevel()
	if level > this.level {
		level = this.level + 1
		this.level = level
	}
	update := make([]*SkipNode, level)
	for i := level - 1; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].num < num {
			p = p.next[i]
		}
		update[i] = p
	}

	// if update[0].next[0] != nil && update[0].next[0].num == num {
	// 	return
	// }

	node := &SkipNode{
		num:  num,
		prev: make([]*SkipNode, level),
		next: make([]*SkipNode, level),
	}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		node.prev[i] = update[i]
		update[i].next[i] = node
	}
}

func (this *Skiplist) Erase(num int) bool {
	p := this.head
	flag := false
	for i := this.level - 1; i >= 0; i-- {
		for p.next[i] != nil {
			if p.next[i].num == num {
				flag = true

				tmp := p.next[i]
				p.next[i] = tmp.next[i]
				tmp.next[i] = nil
				tmp.prev[i] = nil

			} else if p.next[i].num < num {
				p = p.next[i]
			} else {
				break
			}
		}
	}
	if flag {
		for this.level > 1 && this.head.next[this.level-1] == nil {
			this.level--
		}
	}
	return flag
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
