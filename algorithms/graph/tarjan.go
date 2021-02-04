package graph

import "github.com/imkuqin-zw/Data-Structures-and-Algorithms/data_structure/graph"

func Tarjan(g *graph.CFSGraph, p, f int, num *int, dfn, low *[]int) []*graph.CFSEdge {
	result := make([]*graph.CFSEdge, 0)
	*num++
	(*dfn)[p], (*low)[p] = *num, *num
	for i := g.Head[p]; i > 0; i = g.Edge[i].Next {
		if g.Edge[i].To == f {
			continue
		}
		if (*dfn)[g.Edge[i].To] == 0 {
			result = append(result, Tarjan(g, g.Edge[i].To, p, num, dfn, low)...)
			if (*low)[p] > (*low)[g.Edge[i].To] {
				(*low)[p] = (*low)[g.Edge[i].To]
			}
			if (*dfn)[p] < (*low)[g.Edge[i].To] {
				result = append(result, g.Edge[i])
			}
		} else {
			if (*low)[p] > (*dfn)[g.Edge[i].To] {
				(*low)[p] = (*dfn)[g.Edge[i].To]
			}
		}
	}
	return result
}
