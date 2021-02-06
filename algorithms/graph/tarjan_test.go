package graph

import (
	"fmt"
	"testing"

	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/data_structure/graph"
)

func TestUTarjan(t *testing.T) {
	g := graph.NewCFSGraph(7, false)
	g.AddEdge(1, 4, 0)
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 5, 0)
	g.AddEdge(5, 7, 0)
	g.AddEdge(5, 6, 0)
	g.AddEdge(6, 4, 0)

	dfn, low := make([]int, 8), make([]int, 8)
	num := 0
	for i := 1; i <= 7; i++ {
		if dfn[i] == 0 {
			for _, item := range UTarjan(g, i, 0, &num, &dfn, &low) {
				fmt.Println(item)
			}
		}

	}
}

func TestUTarjanP(t *testing.T) {
	g := graph.NewCFSGraph(7, false)
	g.AddEdge(1, 4, 0)
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 5, 0)
	g.AddEdge(5, 7, 0)
	g.AddEdge(5, 6, 0)
	g.AddEdge(6, 4, 0)

	dfn, low := make([]int, 8), make([]int, 8)
	num := 0
	for i := 1; i <= 7; i++ {
		if dfn[i] == 0 {
			for _, item := range UTarjanP(g, i, 0, i, &num, &dfn, &low) {
				fmt.Println(item)
			}
		}
	}
}
