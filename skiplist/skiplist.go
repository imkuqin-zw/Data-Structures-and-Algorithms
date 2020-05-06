package skiplist

//
//import (
//	"math/rand"
//	"time"
//)
//
//const MaxLevel = 32
//const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小
//
//type SkipNode struct {
//	Key  int
//	Val  interface{}
//	prev []*SkipNode
//	next []*SkipNode
//}
//
//type SkipList struct {
//	head  *SkipNode
//	tail  *SkipNode
//	level int
//	len   int
//}
//
//func (sl *SkipList) randLevel() (level int) {
//	rand.Seed(time.Now().UnixNano())
//	level = 1
//	for rand.Float32() < Probability && level < MaxLevel {
//		level++
//	}
//	return
//}
//
//func (sl *SkipList) Get(key int) interface{} {
//	p := sl.head
//	for i := len(p.next) - 1; i >= 0; i-- {
//		for p.next[i] != nil {
//			if p.next[i].Key == key {
//				return p.next[i].Val
//			} else if p.next[i].Key < key {
//				p = p.next[i]
//			} else {
//				break
//			}
//		}
//	}
//	return nil
//}
//
//func (sl *SkipList) Del(key int) {
//	p := sl.head
//	for i := len(p.next) - 1; i >= 0; i-- {
//		for p.next[i] != nil {
//			if p.next[i].Key == key {
//				for {
//					tmp := p.next[i]
//					p.next[i] = tmp.next[i]
//					tmp.next[i] = nil
//					tmp.prev[i] = nil
//					i--
//					if i < 0 {
//						return
//					}
//					p = tmp.prev[i]
//				}
//			} else if p.next[i].Key < key {
//				p = p.next[i]
//			} else {
//				break
//			}
//		}
//	}
//}
//
//func (sl *SkipList) Put(key int, val interface{}) {
//	p := sl.head
//	update := make([]*SkipNode, MaxLevel)
//	for i := sl.level-1; i >= 0; i-- {
//		if p.next[i]!=nil && p.next[i].Key < key {
//			p = p.next[i]
//		}
//		update[i] = p
//	}
//	p = p.next[0]
//	if p != nil && p.Key == key {
//		p.Val = val
//		return
//	}
//
//	level := sl.randLevel()
//	if level > sl.level {
//		level = sl.level + 1
//		update[level]= sl.head
//		sl.level = level
//	}
//	node := &SkipNode{
//		Key:  key,
//		Val:  val,
//		prev: make([]*SkipNode, level),
//		next: make([]*SkipNode, level),
//	}
//	for i := 0; i < level; i++ {
//		node.next[i] = update[i].next[i]
//		node.prev[i] = update[i]
//		update[i].next[i] = node
//	}
//}
