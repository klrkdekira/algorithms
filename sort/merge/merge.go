package merge

type IntSlice []int

func Sort(a []int) []int {
	return MergeSort(a)
}

func Merge(left, right []int) []int {
	var a []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				a = append(a, left[0])
				left = left[1:]
			} else {
				a = append(a, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			a = append(a, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			a = append(a, right[0])
			right = right[1:]
		}
	}
	return a
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return Merge(left, right)
}
