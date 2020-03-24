package sorts

func SelectSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		flag := i
		for j := i + 1; j < length; j++ {
			if data[flag] > data[j] {
				flag = j
			}
		}
		if flag != i {
			data[i], data[flag] = data[flag], data[i]
		}
	}
}

func InsertSort(data []int) {
	if len(data) == 0 {
		return
	}
	for i := 1; i < len(data); i++ {
		tmp, j := data[i], i
		for j >= 1 && data[j-1] > tmp {
			data[j] = data[j-1]
			j--
		}
		if j != i {
			data[j] = tmp
		}
	}
}

func BubbleSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

func MergeSort(data []int, s, e int) {
	merge := func(data []int, s, m, e int) {
		lenL := m - s + 1
		lenR := e - m
		l := make([]int, lenL)
		r := make([]int, lenR)
		copy(l, data[s:m+1])
		copy(r, data[m+1:e+1])
		var i, j, k int
		for k = s; i < lenL && j < lenR; k++ {
			if l[i] <= r[j] {
				data[k] = l[i]
				i++
			} else {
				data[k] = r[j]
				j++
			}
		}
		if i < lenL {
			copy(data[k:e+1], l[i:])
		} else if j < lenR {
			copy(data[k:e+1], r[j:])
		}
	}
	if s < e {
		m := (s + e) / 2
		MergeSort(data, s, m)
		MergeSort(data, m+1, e)
		merge(data, s, m, e)
	}
}

func QuickSort(data []int, low, high int) {
	if low >= high {
		return
	}
	i, j, index := low, high, data[low]
	for i < j {
		for i < j && data[j] >= index {
			j--
		}
		if i < j {
			data[i] = data[j]
			i++
		}
		for i < j && data[i] < index {
			i++
		}
		if i < j {
			data[j] = data[i]
			j--
		}
	}
	data[i] = index
	QuickSort(data, low, i)
	QuickSort(data, i+1, high)
}

func ShellSort(data []int) {
	length := len(data)
	for skip := length / 2; skip > 0; skip /= 2 {
		for i := skip; i < length; i++ {
			j, tmp := i, data[i]
			for j >= skip && data[j-skip] > tmp {
				data[j] = data[j-skip]
				j -= skip
			}
			if j != j-skip {
				data[j] = tmp
			}
		}
	}
}
