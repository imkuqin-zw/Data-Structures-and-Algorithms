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

func merge(data []int, s, m, e int) {
	n1 := m - s + 1
	n2 := e - m
	l := make([]int, n1)
	r := make([]int, n2)
	for i, k := 0, 0; i < n1; i, k = i+1, k+1 {
		l[i] = data[k]
	}
	for i, k := 0, m+1; i < n2; i, k = i+1, k+1 {
		r[i] = data[k]
	}
	var i, j, k int
	for k = s; i < n1 && j < n2; k++ {
		if l[i] <= r[j] {
			data[k] = l[i]
			i++
		} else {
			data[k] = r[j]
			j++
		}
	}
	for ; i < n1; k++ {
		data[k] = l[i]
		i++
	}
	for ; j < n2; k++ {
		data[k] = r[j]
		j++
	}
}

func MergeSort(data []int, s, e int) {

	if s < e {
		m := (s + e) / 2
		MergeSort(data, s, m)
		MergeSort(data, m+1, e)
		merge(data, s, m, e)
	}
}
