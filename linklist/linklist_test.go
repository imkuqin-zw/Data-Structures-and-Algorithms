package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	var l *LinkList
	fmt.Print("traverse: ")
	l.Traverse()
	if result := l.Size(); result != 0 {
		t.Errorf("Size() = %d, want 0", result)
	}
	if result := l.String(); result != "" {
		t.Errorf("string() = %s, want ''", result)
	}
	l = l.InsertAt(0, 3)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l = l.InsertAt(0, 1)
	if result := l.String(); result != "1,3" {
		t.Errorf("string() = %s, want 1,3 ", result)
	}
	l = l.InsertAt(1, 2)
	if result := l.String(); result != "1,2,3" {
		t.Errorf("string() = %s, want 1,2,3 ", result)
	}
	l = l.Remove(2)
	if result := l.String(); result != "1,3" {
		t.Errorf("string() = %s, want 1,3 ", result)
	}
	l = l.Remove(1)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l = l.Remove(0)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l = l.Remove(3)
	if result := l.String(); result != "" {
		t.Errorf("string() = %s, want  ", result)
	}
	l = l.InsertFront(3)
	l = l.InsertFront(2)
	l = l.InsertFront(1)
	l = l.InsertBack(4)
	if result := l.String(); result != "1,2,3,4" {
		t.Errorf("string() = %s, want 1,2,3,4 ", result)
	}
	var val interface{}
	l, val = l.DeleteAt(1)
	if val.(int) != 2 {
		t.Errorf("DeleteAt() val = %s, want 2", val)
	}
	if result := l.String(); result != "1,3,4" {
		t.Errorf("string() = %s, want 1,3,4 ", result)
	}
	l, val = l.DeleteAt(0)
	if val.(int) != 1 {
		t.Errorf("DeleteAt() val = %s, want 1", val)
	}
	if result := l.Size(); result != 2 {
		t.Errorf("Size() = %d, want 2", result)
	}
	l = l.InsertFront(2)
	l = l.InsertFront(1)
	l = l.InsertBack(5)
	l = l.InsertBack(6)
	if result := l.String(); result != "1,2,3,4,5,6" {
		t.Errorf("string() = %s, want 1,2,3,4,5,6 ", result)
	}
	l = l.Reverse()
	if result := l.String(); result != "6,5,4,3,2,1" {
		t.Errorf("string() = %s, want 6,5,4,3 ", result)
	}
	fmt.Print("traverse: ")
	l.Traverse()
}

func TestLinkListH(t *testing.T) {
	var l = new(LinkListH)
	fmt.Print("traverse: ")
	l.Traverse()
	if result := l.Size(); result != 0 {
		t.Errorf("Size() = %d, want 0", result)
	}
	if result := l.String(); result != "" {
		t.Errorf("string() = %s, want ''", result)
	}
	l.InsertAt(0, 3)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l.InsertAt(0, 1)
	if result := l.String(); result != "1,3" {
		t.Errorf("string() = %s, want 1,3 ", result)
	}
	l.InsertAt(1, 2)
	if result := l.String(); result != "1,2,3" {
		t.Errorf("string() = %s, want 1,2,3 ", result)
	}
	l.Remove(2)
	if result := l.String(); result != "1,3" {
		t.Errorf("string() = %s, want 1,3 ", result)
	}
	l.Remove(1)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l.Remove(0)
	if result := l.String(); result != "3" {
		t.Errorf("string() = %s, want 3 ", result)
	}
	l.Remove(3)
	if result := l.String(); result != "" {
		t.Errorf("string() = %s, want  ", result)
	}
	l.InsertFront(3)
	l.InsertFront(2)
	l.InsertFront(1)
	l.InsertBack(4)
	if result := l.String(); result != "1,2,3,4" {
		t.Errorf("string() = %s, want 1,2,3,4 ", result)
	}
	var val interface{}
	val = l.DeleteAt(1)
	if val.(int) != 2 {
		t.Errorf("DeleteAt() val = %s, want 2", val)
	}
	if result := l.String(); result != "1,3,4" {
		t.Errorf("string() = %s, want 1,3,4 ", result)
	}
	val = l.DeleteAt(0)
	if val.(int) != 1 {
		t.Errorf("DeleteAt() val = %s, want 1", val)
	}
	if result := l.Size(); result != 2 {
		t.Errorf("Size() = %d, want 2", result)
	}
	l.InsertFront(2)
	l.InsertFront(1)
	l.InsertBack(5)
	l.InsertBack(6)
	if result := l.String(); result != "1,2,3,4,5,6" {
		t.Errorf("string() = %s, want 1,2,3,4,5,6 ", result)
	}
	l.Reverse()
	if result := l.String(); result != "6,5,4,3,2,1" {
		t.Errorf("string() = %s, want 6,5,4,3 ", result)
	}
	fmt.Print("traverse: ")
	l.Traverse()
}
