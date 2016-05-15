package bubble

import (
	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort"
)

// Found out I NIH the stdlib `sort` XD
func Sort(a []int) []int {
	b := sort.IntSlice(a)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < b.Len()-1; i++ {
			if b.Less(i+1, i) {
				b.Swap(i+1, i)
				swapped = true
			}
		}
	}
	return []int(b)
}

func SortAnimation(a []int, c chan *helper.Progress) {
	b := sort.IntSlice(a)
	c <- &helper.Progress{
		A: []int(b),
		I: -1,
		J: -1,
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < b.Len()-1; i++ {
			if b.Less(i+1, i) {
				b.Swap(i+1, i)
				swapped = true
				c <- &helper.Progress{
					A: []int(b),
					I: i,
					J: i + 1,
				}
			}
		}
	}
	c <- &helper.Progress{
		A: []int(b),
		I: -1,
		J: -1,
	}
	close(c)
}
