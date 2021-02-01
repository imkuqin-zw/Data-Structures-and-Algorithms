package graph

import "testing"

func TestCFSGraph(t *testing.T) {
	g := NewCFSGraph(4, false)
	g.AddEdge(1, 2, 5)
	g.AddEdge(1, 4, 3)
	g.AddEdge(2, 3, 8)
	g.AddEdge(2, 4, 12)
	g.AddEdge(3, 4, 9)
	g.Print()
}
