package bubble_test

import (
	"testing"

	"sort"

	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort/bubble"
)

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 5; i++ {
		a := helper.RandIntN(1e3)
		b := bubble.Sort(a)
		if !sort.IntsAreSorted(b) {
			t.Errorf("Sort result is incorrect, data: %v", b)
		}
	}
}

func benchmarkBubbleSortN(length int, b *testing.B) {
	b.StopTimer()
	a := helper.RandIntN(length)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bubble.Sort(a)
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	benchmarkBubbleSortN(10, b)
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSortN(100, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkBubbleSortN(1000, b)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	benchmarkBubbleSortN(10000, b)
}

func BenchmarkBubbleSort100000(b *testing.B) {
	benchmarkBubbleSortN(100000, b)
}

func BenchmarkBubbleSort1000000(b *testing.B) {
	benchmarkBubbleSortN(1000000, b)
}
