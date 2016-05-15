package insertion

import "github.com/klrkdekira/algorithms/helper"

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

func SortAnimation(a []int, c chan *helper.Progress) {
	c <- &helper.Progress{
		A: []int(a),
		I: -1,
		J: -1,
	}
	for j := 0; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			c <- &helper.Progress{
				A: []int(a),
				I: i,
				J: i + 1,
			}
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	c <- &helper.Progress{
		A: []int(a),
		I: -1,
		J: -1,
	}
	close(c)
}
