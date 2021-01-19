package strings

func kmpNext(substr string) []int {
	next := make([]int, len(substr))
	next[0] = -1
	t := -1
	i := 0
	for i < len(substr)-1 {
		if t == -1 || substr[i] == substr[t] {
			t++
			i++
			if substr[i] != substr[t] {
				next[i] = t
			} else {
				next[i] = next[t]
			}
		} else {
			t = next[t]
		}
	}
	return next
}

func KmpIndex(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	next := kmpNext(substr)
	i := 0
	j := 0
	for i < len(s) && j < len(substr) {
		if j == -1 || s[i] == substr[j] {
			j++
			i++
			if len(substr) == j {
				return i - j
			}
		} else {
			j = next[j]
		}
	}
	return -1
}
