package bubble

import "github.com/klrkdekira/algorithms/sort"

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

func SortAnimation(a []int, c chan []int) {
	b := sort.IntSlice(a)
	c <- []int(b)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < b.Len()-1; i++ {
			if b.Less(i+1, i) {
				b.Swap(i+1, i)
				swapped = true
				c <- []int(b)
			}
		}
	}
	c <- []int(b)
	close(c)
}
