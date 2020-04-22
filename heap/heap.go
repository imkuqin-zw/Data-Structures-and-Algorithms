package heap

type Heap struct {
	elems []interface{}
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

func (h *Heap) DelMax() interface{} {
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

func (h *Heap) Insert(e interface{}) {
	h.elems = append(h.elems, e)
	h.n++
	h.swim(h.n)
}

func (h *Heap) less(i, j int) bool {
	return h.cmp(h.elems[i], h.elems[j]) < 0
}

func (h *Heap) exchange(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

// 上浮
func (h *Heap) swim(i int) {
	j := i / 2
	for j > 0 && h.less(j, i) {
		h.exchange(i, j)
		i = j
		j = i / 2
	}
}

// 下沉
func (h *Heap) sink(i int, n int) {
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

// 排序
func (h *Heap) Sort() []interface{} {
	i := h.n
	for i > 1 {
		h.exchange(1, i)
		i--
		h.sink(1, i)
	}
	return h.elems[1:]
}
