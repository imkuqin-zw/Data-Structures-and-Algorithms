package binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMaxSubTree(t *testing.T) {
	type args struct {
		root *BNode
	}
	type testParams struct {
		name  string
		arg   args
		want  int
		want1 *BNode
	}
	type createBTree func([]interface{}, int, int, int, **BNode) *BNode
	var create createBTree
	create = func(data []interface{}, start, end int, wantIndex int, want1 **BNode) *BNode {
		var root *BNode
		if start <= end {
			mid := (start + end + 1) / 2
			root = &BNode{
				Data:   data[mid],
				LChild: create(data, start, mid-1, wantIndex, want1),
				RChild: create(data, mid+1, end, wantIndex, want1),
			}
			if mid == wantIndex {
				*want1 = root
			}
		}
		return root
	}
	fn := func(vals []interface{}, want int, wantIndex int) testParams {
		var params testParams
		params.arg = args{create(vals, 0, len(vals)-1, wantIndex, &params.want1)}
		params.want = want
		return params
	}
	var tests = []testParams{
		fn([]interface{}{-1, 3, 6, -7}, 2, 1),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindMaxSubTree(tt.arg.root)
			if got != tt.want {
				t.Errorf("FindMaxSubTree() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindMaxSubTree() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBTreeToDLink(t *testing.T) {
	b := (*BNode)(CreateBTree([]interface{}{1, 2, 3, 4, 5, 6, 7}, 0, 6))
	dl := BTreeToDLink(b)
	cur := dl
	var last = cur
	for cur != nil {
		fmt.Print(cur.Data, " ")
		if cur.RChild != nil {
			last = cur.RChild
		}
		cur = cur.RChild
	}
	fmt.Println()
	for last != nil {
		fmt.Print(last.Data, " ")
		last = last.LChild
	}
	fmt.Println()
}
