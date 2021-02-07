package graph

import "fmt"

type AMGraph struct {
	Vex      []interface{}
	Edge     [][]int
	EdgeNum  int
	Directed bool
}

func NewAMGraph(vex []interface{}, directed bool) *AMGraph {
	g := &AMGraph{
		Directed: directed,
		Vex:      vex,
		Edge:     make([][]int, len(vex)),
	}
	for i := 0; i < len(vex); i++ {
		g.Edge[i] = make([]int, len(vex))
	}
	return g
}

func (g *AMGraph) FindVex(vex interface{}) int {
	for i := 0; i < len(g.Vex); i++ {
		if g.Vex[i] == vex {
			return i
		}
	}
	return -1
}

func (g *AMGraph) AddEdge(from, to interface{}, val int) {
	i := g.FindVex(from)
	if i == -1 {
		return
	}
	j := g.FindVex(to)
	if i == -1 {
		return
	}
	g.Edge[i][j] = val
	if !g.Directed {
		g.Edge[j][i] = val
		g.EdgeNum++
	}
	g.EdgeNum++
}

func (g *AMGraph) Print() {
	for i := 0; i < len(g.Edge); i++ {
		for j := 0; j < len(g.Edge[i]); j++ {
			fmt.Printf("%d\t", g.Edge[i][j])
		}
		fmt.Println()
	}
}
