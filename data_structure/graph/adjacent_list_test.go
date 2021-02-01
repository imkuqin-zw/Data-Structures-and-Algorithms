package graph

import "testing"

func TestALGraph(t *testing.T) {
	g := NewALGraph([]ALVexType{'a', 'b', 'c', 'd', 'e'}, true)
	g.AddEdge('a', 'b', 1)
	g.AddEdge('a', 'c', 1)
	g.AddEdge('a', 'e', 1)
	g.AddEdge('b', 'c', 1)
	g.AddEdge('c', 'd', 1)
	g.AddEdge('c', 'e', 1)
	g.AddEdge('d', 'e', 1)
	g.Print()
}
