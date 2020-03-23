package BTree

import (
	"fmt"
	"testing"
)

func TestWordDictionary_AddWord(t *testing.T) {
	words := []string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search", "search", "search"}
	wt := WordDictionary{Children: make([]*WordDictionary, 0)}
	for _, item := range words {
		wt.AddWord(item)
	}
	fmt.Println(wt.Search(".a"))
}
