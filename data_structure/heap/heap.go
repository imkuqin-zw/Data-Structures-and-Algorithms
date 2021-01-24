package heap

type Heap struct {
	val     []int
	greater bool
}

func NewHeap(greater bool) *Heap {
	return &Heap{
		val:     make([]int, 1),
		greater: greater,
	}
}

func (h *Heap) Push(val int) {
	h.val = append(h.val, val)
	h.swim(len(h.val) - 1)
}

func (h *Heap) Pop() (bool, int) {
	if len(h.val) <= 1 {
		return false, 0
	}
	val := h.val[1]
	h.val[1] = h.val[len(h.val)-1]
	h.val = h.val[:len(h.val)-1]
	h.sink(1)
	return true, val
}

func (h *Heap) needExch(c, p int) bool {
	if h.greater {
		return h.val[c] > h.val[p]
	} else {
		return h.val[c] < h.val[p]
	}
}

func (h *Heap) swim(idx int) {
	for idx > 1 {
		pIdx := idx / 2
		if h.needExch(idx, pIdx) {
			h.val[pIdx], h.val[idx] = h.val[idx], h.val[pIdx]
		}
		idx = pIdx
	}
}

func (h *Heap) sink(idx int) {
	for 2*idx < len(h.val) {
		nIdx := idx * 2
		rc := idx*2 + 1
		if rc < len(h.val)-1 && h.needExch(rc, nIdx) {
			nIdx = rc
		}
		if h.needExch(nIdx, idx) {
			h.val[nIdx], h.val[idx] = h.val[idx], h.val[nIdx]
		}
		idx = nIdx
	}
}
