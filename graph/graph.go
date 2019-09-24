package graph

import "math"

type Graph struct {
}

type VStatus int8

const (
	_ VStatus = iota
	Undiscovered
	Discovered
	Visited
)

//顶点
type Vertex struct {
	data                interface{}
	inDegree, outDegree int //入度、出度
	status              VStatus
	dTime, fTime        int //时间标签
	parent              int //在遍历树中的父节点
	priority            int //在遍历树中的优先级（最短通路、极短跨边等）
}

func NewVertex(data interface{}) *Vertex {
	return &Vertex{
		data:     data,
		status:   Undiscovered,
		fTime:    -1,
		dTime:    -1,
		parent:   -1,
		priority: math.MaxInt32,
	}
}

type EStatus int

const (
	_ EStatus = iota
	Undetermined
	Tree
	Cross
	Forward
	Backward
)

type Edge struct {
	Data   interface{}
	weight int
	status EStatus
}

func NewEdge(weight int) *Edge {
	return &Edge{weight: weight, status: Undetermined}
}

//func NewEdge(data interface{}, weight int) *Edge {
//	return &Edge{
//		Data:   data,
//		weight: weight,
//		status: Undetermined,
//	}
//}
