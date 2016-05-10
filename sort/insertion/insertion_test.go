package insertion_test

import (
	"testing"

	"sort"

	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort/insertion"
)

func TestInsertionSort(t *testing.T) {
	for i := 0; i < 5; i++ {
		a := insertion.Sort(helper.RandIntN(1e3))
		if !sort.IntsAreSorted(a) {
			t.Errorf("Sort result is incorrect, data: %v", a)
		}
	}
}

func benchmarkInsertionSortN(length int, b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		a := helper.RandIntN(length)
		b.StartTimer()
		insertion.Sort(a)
		b.StopTimer()
	}
}

func BenchmarkInsertionSort10(b *testing.B) {
	benchmarkInsertionSortN(10, b)
}

func BenchmarkInsertionSort100(b *testing.B) {
	benchmarkInsertionSortN(100, b)
}

func BenchmarkInsertionSort1000(b *testing.B) {
	benchmarkInsertionSortN(1000, b)
}

func BenchmarkInsertionSort10000(b *testing.B) {
	benchmarkInsertionSortN(10000, b)
}

func BenchmarkInsertionSort100000(b *testing.B) {
	benchmarkInsertionSortN(100000, b)
}

func BenchmarkInsertionSort1000000(b *testing.B) {
	benchmarkInsertionSortN(1000000, b)
}
