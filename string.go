package main

func getNext(sub string) []int {
	subLen := len(sub)
	next := make([]int, subLen)
	next[0] = -1
	next[1] = 0
	for i := 2; i < subLen; i++ {
		next[i] = getCurNext(sub, i)
	}
	return next
}

func getCurNext(sub string, cur int) int {
	matched := 0
	right := cur - 1
	left := (right - 1) / 2
	curLeft := left
	curRight := right
	for curLeft >= 0 {
		if sub[curLeft] == sub[curRight] {
			curLeft--
			curRight--
			matched++
		} else {
			matched = 0
			left--
			curLeft = left
			curRight = right
		}
	}
	return matched
}

func index(src, sub string) int {
	if sub == "" {
		return -1
	}
	next := map[int]int{0: -1, 1: 0} //getNext(sub)
	srcLen := len(src)
	subLen := len(sub)
	var j, i = 0, 0
	for i < srcLen && j < subLen {
		if j == -1 || src[i] == sub[j] {
			i++
			j++
		} else {
			// BF
			// i = i - j + 1
			// j = 0

			// KMP
			if _, ok := next[j]; !ok {
				next[j] = getCurNext(sub, j)
			}
			j = next[j]
		}

	}
	if j < subLen {
		return -1
	}
	return i - subLen
}

func commStr(a, b string) string {
	aLen, bLen := len(a), len(b)
	if aLen == 0 || bLen == 0 {
		return ""
	}
	aCur, bCur := aLen-1, bLen-1
	if a[aCur] == b[bCur] {
		return commStr(a[0:aCur], b[0:bCur]) + string(a[aCur])
	}
	aSub := commStr(a, b[0:bCur])
	bSub := commStr(a[0:aCur], b)
	if len(aSub) >= len(bSub) {
		return aSub
	}
	return bSub
}
