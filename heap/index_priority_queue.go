package heap

import "container/list"

type IndexPriorityQueue struct {
	elems []interface{}
	pq    []int
	qp    []int
	n     int
	cmp   func(a, b interface{}) int
}

func New(cmp func(a, b interface{}) int, capacity int) *Heap {
	capacity++
	return &Heap{
		elems: make([]interface{}, 1, capacity),
		cmp:   cmp,
	}
}

func (h *IndexPriorityQueue) Pop() interface{} {
	if h.n == 0 {
		return nil
	}
	max := h.elems[1]
	h.elems[1] = h.elems[h.n]
	h.elems = h.elems[0:h.n]
	h.n--
	h.sink(1, h.n)
	return max
}

func (h *IndexPriorityQueue) Push(e interface{}) int {
	h.elems = append(h.elems, e)
	h.pq = append(h.pq, h.n)
	index := h.n
	h.n++
	h.swim(h.n)
	return index
}

// 排序
func (h *IndexPriorityQueue) Sort() []interface{} {
	i := h.n
	for i > 1 {
		h.exchange(1, i)
		i--
		h.sink(1, i)
	}
	return h.elems[1:]
}

func (h *IndexPriorityQueue) IsEmpty() bool {
	return h.n == 0
}

func (h *IndexPriorityQueue) Size() int {
	return h.n
}

func (h *IndexPriorityQueue) less(i, j int) bool {
	return h.cmp(h.elems[i], h.elems[j]) < 0
}

func (h *IndexPriorityQueue) exchange(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

// 上浮
func (h *IndexPriorityQueue) swim(i int) {
	j := i / 2
	for j > 0 && h.less(j, i) {
		h.exchange(i, j)
		i = j
		j = i / 2
	}
}

// 下沉
func (h *IndexPriorityQueue) sink(i int, n int) {
	for l := 2 * i; l <= n; l = 2 * i {
		max := l
		r := l + 1
		if r <= n && h.less(l, r) {
			max = r
		}
		if h.less(max, i) {
			return
		}
		h.exchange(i, max)
		i = max
	}
}

func mySqrt(x int) int {
	l := 0
	r := x
	result := -1
	for l <= r {
		mid := (r + l) / 2
		if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}

func numIslands(grid [][]byte) int {
	type point struct {
		x, y int
	}
	num := 0
	q := list.New()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				num++
				grid[i][j] = 0
				q.PushBack(&point{i, j})
				for q.Len() > 0 {
					e := q.Front()
					q.Remove(e)
					p := e.Value.(*point)
					if p.x+1 < len(grid) && grid[p.x+1][p.y] == 1 {
						grid[p.x+1][p.y] = 0
						q.PushBack(&point{p.x + 1, p.y})
					}
					if p.x-1 >= 0 && grid[p.x-1][p.y] == 1 {
						grid[p.x-1][p.y] = 0
						q.PushBack(&point{p.x - 1, p.y})
					}
					if p.y+1 < len(grid[p.x]) && grid[p.x][p.y+1] == 1 {
						grid[p.x][p.y+1] = 0
						q.PushBack(&point{p.x, p.y + 1})
					}
					if p.y-1 >= 0 && grid[p.x][p.y-1] == 1 {
						grid[p.x][p.y-1] = 0
						q.PushBack(&point{p.x, p.y - 1})
					}
				}
			}
		}
	}
	return num
}
