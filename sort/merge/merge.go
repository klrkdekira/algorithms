package merge

import "github.com/klrkdekira/algorithms/helper"

type (
	IntSlice []int

	Merger struct{}
)

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

func (s *Merger) SortAnimation(a []int, c chan *helper.Progress) {
	c <- &helper.Progress{
		A: a,
		I: -1,
		J: -1,
	}
	MergeSortAnimation(a, 0, len(a), c)
	c <- &helper.Progress{
		A: a,
		I: -1,
		J: -1,
	}
	close(c)
}

func MergeSortAnimation(a []int, p, r int, c chan *helper.Progress) {
	if r-p < 2 {
		return
	}

	if p < r {
		q := (p + r) / 2
		MergeSortAnimation(a, p, q, c)
		MergeSortAnimation(a, q, r, c)
		MergeAnimation(a, p, q, r, c)
	}
}

func MergeAnimation(a []int, p, q, r int, c chan *helper.Progress) {
	n1 := q - p
	n2 := r - q

	left := make([]int, n1)
	for i := 0; i < n1; i++ {
		left[i] = a[p+i]
	}

	right := make([]int, n2)
	for j := 0; j < n2; j++ {
		right[j] = a[q+j]
	}

	i, j := 0, 0
	for k := p; k < r; k++ {
		if len(left) > i || len(right) > j {
			if len(left) > i && len(right) > j {
				if left[i] <= right[j] {
					a[k] = left[i]
					c <- &helper.Progress{
						A: a,
						I: i,
						J: k,
					}
					i++
				} else {
					a[k] = right[j]
					c <- &helper.Progress{
						A: a,
						I: k,
						J: j,
					}
					j++
				}
			} else if len(left) > i {
				a[k] = left[i]
				c <- &helper.Progress{
					A: a,
					I: i,
					J: k,
				}
				i++
			} else if len(right) > j {
				a[k] = right[j]
				c <- &helper.Progress{
					A: a,
					I: k,
					J: j,
				}
				j++
			}

		}
	}
}
