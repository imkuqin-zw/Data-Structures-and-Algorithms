package tree

import (
	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/data_structure/heap"
)

/**
总时间限制: 1000ms 内存限制: 65536kB
描述
构造一个具有n个外部节点的扩充二叉树，每个外部节点Ki有一个Wi对应，作为该外部节点的权。使得这个扩充二叉树的叶节点带权外部路径长度总和最小：

									 Min( W1 * L1 + W2 * L2 + W3 * L3 + … + Wn * Ln)

Wi:每个节点的权值。

Li:根节点到第i个外部叶子节点的距离。

编程计算最小外部路径长度总和。

输入
第一行输入一个整数n，外部节点的个数。第二行输入n个整数，代表各个外部节点的权值。
2<=N<=100
输出
输出最小外部路径长度总和。
样例输入
4
1 1 3 5
样例输出
17
*/
func BaiLian4080(val []int) int {
	if len(val) < 2 {
		if len(val) == 1 {
			return val[0]
		}
		return 0
	}
	h := heap.NewHeap(false)
	for _, item := range val {
		h.Push(item)
	}
	result := 0
	for {
		b, l := h.Pop()
		if !b {
			break
		}
		b, r := h.Pop()
		if !b {
			break
		}
		result += r + l
		h.Push(r + l)
	}
	return result
}

/**
https://www.luogu.com.cn/problem/UVA240
*/
