package heap

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
