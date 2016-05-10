package insertion

func Sort(a []int) []int {
	return SortWithLength(a, len(a))
}

func SortWithLength(a []int, length int) []int {
	if length == 0 {
		return []int{}
	}

	if length == 1 {
		return a
	}

	for j := 0; j < length; j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return a
}
