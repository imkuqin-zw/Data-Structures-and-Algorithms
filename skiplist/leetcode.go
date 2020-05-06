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

//["Skiplist","add","add","add","add","add","add","add","add","add","erase","search","add","erase","erase","erase","add","search","search","search","erase","search","add","add","add","erase","search","add","search","erase","search","search","erase","erase","add","erase","search","erase","erase","search","add","add","erase","erase","erase","add","erase","add","erase","erase","add","add","add","search","search","add","erase","search","add","add","search","add","search","erase","erase","search","search","erase","search","add","erase","search","erase","search","erase","erase","search","search","add","add","add","add","search","search","search","search","search","search","search","search","search"]
//70
//[null,null,null,null,null,null,null,null,null,null,true,false,null,true,false,false,null,true,true,true,true,false,null,null,null,false,false,null,false,false,true,true,false,false,null,true,true,false,true,true,null,null,false,true,false,null,true,null,true,true,null,null,null,false,false,null,true,false,null,null,true,null,false,false,false,true,true,false,false,null,true,false,false,false,false,true,false,false,null,null,null,null,true,true,true,true,true,true,false,false,true]
//[null,null,null,null,null,null,null,null,null,null,true,false,null,true,false,false,null,true,true,true,true,false,null,null,null,false,false,null,false,false,true,true,false,false,null,true,true,false,true,true,null,null,false,true,false,null,true,null,true,true,null,null,null,false,false,null,true,false,null,null,true,null,false,false,false,true,true,false,true,null,true,false,false,false,true,true,false,false,null,null,null,null,true,true,true,true,true,true,false,false,true]

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
