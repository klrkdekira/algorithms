package helper

import (
	"math/rand"
	"time"
)

type Progress struct {
	A []int
	I int
	J int
}

func RandIntN(length int) []int {
	a := make([]int, length)
	for i := 0; i < length; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		a[i] = rand.Intn(length)
	}
	return a
}
