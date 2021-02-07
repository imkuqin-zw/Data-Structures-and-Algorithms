package graph

import (
	"math"

	"github.com/imkuqin-zw/Data-Structures-and-Algorithms/data_structure/graph"
)

func Dijkstra(g *graph.AMGraph, u int) {
	var dist = g.Edge[u]                //make([]int, len(g.Vex))  //最短路径长度
	var use = make([]bool, len(g.Vex))  //已经找到的节点
	var front = make([]int, len(g.Vex)) //到达各个顶点的上一个顶点

	for i := 0; i < len(g.Vex); i++ {
		min, t := math.MaxInt32, u
		for j := 0; j < len(g.Vex); j++ {
			if !use[j] && dist[j] > min {
				min = dist[j]
				t = j
			}
		}
		if u == t {
			return
		}
		use[t] = true
		for j := 0; j < len(g.Vex); j++ {
			if !use[j] && g.Edge[t][j] != math.MaxInt32 {
				if dist[j] > dist[t]+g.Edge[t][j] {
					dist[j] = dist[t] + g.Edge[t][j]
					front[j] = t
				}
			}
		}
	}
}

func Floyd(g *graph.AMGraph) {
	var dist = make([][]int, len(g.Vex))  //最短路径长度
	var front = make([][]int, len(g.Vex)) //到达各个顶点的上一个顶点

	for i := 0; i < len(g.Vex); i++ {
		for j := 0; j < len(g.Vex); j++ {
			dist[i][j] = g.Edge[i][j]
			if dist[i][j] < math.MaxInt32 && i != j {
				front[i][j] = i
			} else {
				front[i][j] = -1
			}
		}
	}

	for k := 0; k < len(g.Vex); k++ {
		for i := 0; i < len(g.Vex); i++ {
			for j := 0; j < len(g.Vex); j++ {
				length := dist[i][k] + dist[k][j]
				if dist[i][j] > length {
					dist[i][j] = length
					front[i][j] = front[k][j]
				}
			}
		}
	}
}
