package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	Size = 75
)

func BenchmarkLinearSearch(b *testing.B) {
	b.StopTimer()

	lists := make([]int, 0)
	for i := 0; i < Size; i++ {
		lists = append(lists, i)
	}
	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		target := rand.Intn(len(lists))
		for _, v := range lists {
			if target == v {
				break
			}
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	b.StopTimer()

	lists := make([]int, 0)
	for i := 0; i < Size; i++ {
		lists = append(lists, i)
	}
	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		start, end := 0, len(lists)-1
		target := rand.Intn(len(lists))
		for start <= end {
			pivot := (start + end) / 2
			if lists[pivot] < target {
				start = pivot + 1
			} else if lists[pivot] > target {
				end = pivot - 1
			} else {
				break
			}
		}
	}
}
