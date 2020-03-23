package graph

import "github.com/imkuqin-zw/Data-Structures-and-Algorithms/queue/arrayQueue"

type Matrix struct {
	v []*Vertex
	e [][]*Edge
	n int
}

func (m *Matrix) GetVertex(i int) interface{} {
	return m.v[i].data
}

func (m *Matrix) GetInDegree(i int) int {
	return m.v[i].inDegree
}

func (m *Matrix) GetOutDegree(i int) int {
	return m.v[i].outDegree
}

func (m *Matrix) GetStatus(i int) VStatus {
	return m.v[i].status
}

func (m *Matrix) GetDTime(i int) int {
	return m.v[i].dTime
}

func (m *Matrix) GetFTime(i int) int {
	return m.v[i].fTime
}

func (m *Matrix) GetParent(i int) int {
	return m.v[i].parent
}

func (m *Matrix) GetPriority(i int) int {
	return m.v[i].priority
}

func (m *Matrix) SetDTime(i, clock int) {
	m.v[i].dTime = clock
}

func (m *Matrix) SetVStatus(i int, status VStatus) {
	m.v[i].status = status
}

func (m *Matrix) GetVStatus(i int) VStatus {
	return m.v[i].status
}

func (m *Matrix) SetEStatus(i, j int, status EStatus) {
	m.e[i][j].status = status
}

func (m *Matrix) SetParent(i, parent int) {
	m.v[i].parent = parent
}

func (m *Matrix) nextNbr(i, j int) int {
	for -1 < j {
		j--
		if m.Exist(i, j) {
			break
		}
	}
	return j
}

func (m *Matrix) FirstBbr(i int) int {
	return m.nextNbr(i, m.n)
}

func (m *Matrix) Exist(i, j int) bool {
	return (0 <= i && i < m.n) && (0 <= j && j < m.n) && m.e[i][j] != nil
}

// 插入边
func (m *Matrix) EdgeInsert(i, j, w int) {
	if m.Exist(i, j) {
		return
	}
	m.e[i][j] = NewEdge(w)
	m.v[i].outDegree++
	m.v[j].inDegree++
}

// 删除边
func (m *Matrix) EdgeRemove(i, j int) *Edge {
	if !m.Exist(i, j) {
		return nil
	}
	eBak := m.e[i][j]
	m.e[i][j] = nil
	m.v[i].outDegree--
	m.v[j].inDegree--
	return eBak
}

// 插入顶点
func (m *Matrix) VertexInsert(data interface{}) {
	for i := 0; i < m.n; i++ {
		m.e[i] = append(m.e[i], nil)
	}
	m.n++
	nEdge := make([]*Edge, m.n)
	m.e = append(m.e, nEdge)
	m.v = append(m.v, NewVertex(data))
}

// 删除顶点
func (m *Matrix) VertexRemove(i int) *Vertex {
	for j := 0; j < m.n; j++ {
		if m.Exist(i, j) {
			m.v[j].inDegree--
		}
	}
	m.e = append(m.e[:i], m.e[i+1:]...)
	m.n--
	for j := 0; j < m.n; j++ {
		m.v[j].inDegree--
		m.e[j] = append(m.e[j][:i], m.e[j][i+1:]...)
	}
	vBak := m.v[i]
	m.v = append(m.v[:i], m.v[i+1:]...)
	return vBak
}

// 广度优先遍历
func (m *Matrix) bfs(i, clock int) {
	que := arrayQueue.NewQueue()
	m.SetVStatus(i, Discovered)
	que.Enqueue(i)
	for !que.IsEmpty() {
		v := que.Dequeue().(int)
		clock++
		m.SetDTime(v, clock)
		for u := m.FirstBbr(v); -1 < u; u = m.nextNbr(v, u) {
			if m.GetVStatus(u) == Undiscovered {
				m.SetVStatus(u, Discovered)
				que.Enqueue(u)
				m.SetEStatus(v, u, Tree)
				m.SetParent(u, v)
			} else {
				m.SetEStatus(v, u, Cross)
			}
		}
		m.SetVStatus(i, Visited)
	}
}

func (m *Matrix) BFS(i int) {
	clock := 0
	v := i
	for {
		if m.GetStatus(v) == Discovered {
			m.bfs(v, clock)
		}
		v = (v + 1) % m.n
		if v != i {
			break
		}
	}
}
