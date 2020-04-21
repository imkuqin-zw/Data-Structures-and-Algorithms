package heap

type Heap struct {
	elems []interface{}
	n     int
	cmp   func(a, b interface{}) int
}

func New(cmp func(a, b interface{}) int, capacity int) *Heap {
	return &Heap{
		elems: make([]interface{}, 1, capacity),
		cmp:   cmp,
	}
}

func (h *Heap) Less(i, j int) bool {
	return h.cmp(h.elems[i], h.elems[j]) < 0
}

func (h *Heap) Exchange(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *Heap) DelMax() {
	if h.n == 0 {
		return
	}
	if h.n == 1 {
		h.elems = h.elems[0:1]
		return
	}
	h.elems[1] = h.elems[h.n]
	h.elems = h.elems[0:h.n]
	h.n--


}

func (h *Heap) Insert(e interface{}) {
	h.elems = append(h.elems, e)
	h.n++
	h.Swim()
}

// 上浮
func (h *Heap) Swim() {
	i, j := h.n/2, h.n
	for i > 0 && h.Less(i, j) {
		h.Exchange(i, j)
		j = i
		i = j / 2
	}
}

// 下沉
func (h *Heap) Sink() {
	i := 1
	l, r := 2*i, 2*i+1
	for l <= h.n {
		max := l
		if r <= h.n && h.Less(l, r){
			max = r
		}
		if h.Less(i, max) {
			h.Exchange(i, max)
		}
	}
}
