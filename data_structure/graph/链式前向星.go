package graph

import "fmt"

type CFSEdge struct {
	To, Next, W int
}

type CFSGraph struct {
	Directed bool
	Head     []int
	Edge     []*CFSEdge
}

func NewCFSGraph(vexNum int, directed bool) *CFSGraph {
	g := &CFSGraph{
		Directed: directed,
		Head:     make([]int, vexNum+1),
		Edge:     make([]*CFSEdge, 1),
	}
	return g
}

func (g *CFSGraph) addEdge(f, t, w int) {
	e := &CFSEdge{To: t, W: w}
	g.Edge = append(g.Edge, e)
	idx := len(g.Edge) - 1
	e.Next = g.Head[f]
	g.Head[f] = idx
}

func (g *CFSGraph) AddEdge(from, to, weight int) {
	g.addEdge(from, to, weight)
	if !g.Directed {
		g.addEdge(to, from, weight)
	}
}

func (g *CFSGraph) Print() {
	for i := 1; i < len(g.Head); i++ {
		fmt.Printf("%2d", i)
		j := g.Head[i]
		for j != 0 {
			fmt.Printf("\t%2d:(%2d %2d %2d)", j, g.Edge[j].To, g.Edge[j].W, g.Edge[j].Next)
			j = g.Edge[j].Next
		}
		fmt.Println()
	}
}
