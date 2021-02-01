package graph

import "fmt"

type (
	ALVexType  byte
	ALEdgeType int
)

type ALVexNode struct {
	Val   ALVexType
	First *ALAdjNode
}

type ALAdjNode struct {
	V    int
	W    ALEdgeType
	Next *ALAdjNode
}

type ALGraph struct {
	Directed bool
	Head     []*ALVexNode
	EdgeNum  int
}

func NewALGraph(vex []ALVexType, directed bool) *ALGraph {
	g := &ALGraph{
		Directed: directed,
		Head:     make([]*ALVexNode, len(vex)),
	}
	for i := 0; i < len(vex); i++ {
		g.Head[i] = &ALVexNode{Val: vex[i]}
	}
	return g
}

func (g *ALGraph) findNode(vex ALVexType) int {
	for i, v := range g.Head {
		if v.Val == vex {
			return i
		}
	}
	return -1
}

func (g *ALGraph) addEdge(from, to ALVexType, weight ALEdgeType) {
	i := g.findNode(from)
	if i == -1 {
		return
	}
	j := g.findNode(to)
	if i == -1 {
		return
	}
	g.Head[i].First = &ALAdjNode{
		V:    j,
		W:    weight,
		Next: g.Head[i].First,
	}
	g.EdgeNum++
}

func (g *ALGraph) AddEdge(from, to ALVexType, weight ALEdgeType) {
	g.addEdge(from, to, weight)
	if !g.Directed {
		g.addEdge(to, from, weight)
	}
}

func (g *ALGraph) Print() {
	for i := 1; i < len(g.Head); i++ {
		fmt.Printf("%2d: %c", i, g.Head[i].Val)
		next := g.Head[i].First
		for next != nil {
			fmt.Printf("\t%2d:(%2c %2c %2d)", next.V, g.Head[i].Val, g.Head[next.V].Val, next.W)
			next = next.Next
		}
		fmt.Println()
	}
}
