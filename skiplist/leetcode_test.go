package skiplist

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	sl := Constructor()
	sl.println()
	sl.Add(1)
	sl.println()
	sl.Add(2)
	sl.println()
	sl.Add(3)
	sl.println()
	sl.Add(4)
	sl.println()
	sl.Add(1)
	sl.println()
	sl.Add(4)
	sl.println()
	//fmt.Println(sl.Search(1))
	//fmt.Println(sl.Erase(1))
	//sl.println()
	fmt.Println(sl.Erase(3))
	sl.println()
	fmt.Println(sl.Erase(4))
	sl.println()
	fmt.Println(sl.Erase(1))
	sl.println()
	fmt.Println(sl.Erase(2))
	sl.println()
	fmt.Println(sl.Erase(2))
	sl.println()
	fmt.Println(sl.Erase(1))
	//fmt.Println(sl.Search(2))
	//fmt.Println(sl.Search(1))
	//fmt.Println(sl.Search(3))
	//fmt.Println(sl.Erase(3))
	//fmt.Println(sl.Search(9))
}
