package graph

type Matrix struct {
	v []*Vertex
	e [][]*Edge
	n int
}

func (m *Matrix) Vertex(i int) interface{} {
	return m.v[i].data
}

func (m *Matrix) InDegree(i int) int {
	return m.v[i].inDegree
}

func (m *Matrix) OutDegree(i int) int {
	return m.v[i].outDegree
}

func (m *Matrix) Status(i int) VStatus {
	return m.v[i].status
}

func (m *Matrix) DTime(i int) int {
	return m.v[i].dTime
}

func (m *Matrix) FTime(i int) int {
	return m.v[i].fTime
}

func (m *Matrix) Parent(i int) int {
	return m.v[i].parent
}

func (m *Matrix) priority(i int) int {
	return m.v[i].priority
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
