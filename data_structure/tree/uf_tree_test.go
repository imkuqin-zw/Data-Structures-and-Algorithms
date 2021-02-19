package tree

import (
	"fmt"
	"testing"
)

func TestUFTree(t *testing.T) {
	uf := NewUFTree(5)
	fmt.Println(uf.Count())
	uf.Union(0, 1)
	fmt.Println(uf.Count(), uf.Connected(0, 1))
	uf.Union(0, 2)
	fmt.Println(uf.Count(), uf.Connected(1, 2))
	uf.Union(2, 3)
	fmt.Println(uf.Count(), uf.Connected(0, 2))
	uf.Union(1, 4)
	fmt.Println(uf.Count(), uf.Connected(0, 2))

	uf.Union(1, 2)
	fmt.Println(uf.Count(), uf.Connected(0, 2))
}
