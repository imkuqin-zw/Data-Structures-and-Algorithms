package strings

func badChar(substr string) map[byte]int {
	bc := make(map[byte]int)
	for i, b := range substr {
		bc[byte(b)] = i
	}
	return bc
}

func goodSuffix(substr string) (map[int]int, map[int]struct{}) {
	suffix := make(map[int]int)
	prefix := make(map[int]struct{})
	for i := 0; i < len(substr)-1; i++ {
		j := i
		k := 0
		for j >= 0 && substr[j] == substr[len(substr)-k-1] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j < 0 {
			prefix[k] = struct{}{}
		}
	}
	return suffix, prefix
}

func gsMove(j, subLen int, suffix map[int]int, prefix map[int]struct{}) int {
	k := subLen - j - 1
	if y, ok := suffix[k]; ok {
		return j - y + 1
	}

	for k = k - 1; k > 0; k-- {
		if _, ok := prefix[k]; ok {
			return j - k
		}
	}
	return subLen
}

func bcMove(j int, p byte, bc map[byte]int) int {
	if x, ok := bc[p]; ok {
		return j - x
	}
	return j + 1
}

func BmIndex(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	bc := badChar(substr)
	suffix, prefix := goodSuffix(substr)
	i := 0
	for i <= len(s)-len(substr) {
		j := len(substr) - 1
		for j >= 0 && s[i+j] == substr[j] {
			j--
		}
		if j < 0 {
			return i
		}
		x := bcMove(j, s[j+i], bc)
		var y int
		if len(substr)-1 > 0 && j < len(substr)-1 {
			y = gsMove(j, len(substr), suffix, prefix)
		}
		if y > x {
			i += y
		} else {
			i += x
		}
	}
	return -1
}
