package merge_test

import (
	"sort"
	"testing"

	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort/merge"
)

func TestMergeSort(t *testing.T) {
	for i := 0; i < 5; i++ {
		a := merge.Sort(helper.RandIntN(1e3))
		if !sort.IntsAreSorted(a) {
			t.Errorf("Sort result is incorrect, data: %v", a)
		}
	}
}

func benchmarkMergeSortN(length int, b *testing.B) {
	b.StopTimer()
	a := helper.RandIntN(length)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		merge.Sort(a)
	}
}

func BenchmarkMergeSort10(b *testing.B) {
	benchmarkMergeSortN(10, b)
}

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSortN(100, b)
}

func BenchmarkMergeSort1000(b *testing.B) {
	benchmarkMergeSortN(1000, b)
}

func BenchmarkMergeSort10000(b *testing.B) {
	benchmarkMergeSortN(10000, b)
}

func BenchmarkMergeSort100000(b *testing.B) {
	benchmarkMergeSortN(100000, b)
}

func BenchmarkMergeSort1000000(b *testing.B) {
	benchmarkMergeSortN(1000000, b)
}
