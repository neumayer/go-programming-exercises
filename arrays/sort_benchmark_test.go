package arrays

import (
	"math/rand"
	sortg "sort"
	"testing"
)

var bs = bubblesort{}
var qs = quicksort{}
var res = []int{}

type randomIntSlice struct {
	values []int
}

func (ris *randomIntSlice) getCopy() []int {
	copiedValues := make([]int, len(ris.values))
	copy(copiedValues, ris.values)
	return copiedValues
}

var intSlice100 = randomIntSlice{makeRandomIntSlice(100)}
var intSlice1000 = randomIntSlice{makeRandomIntSlice(1000)}
var intSlice10000 = randomIntSlice{makeRandomIntSlice(10000)}

func makeRandomIntSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func benchmarkSort(s sort, data randomIntSlice, b *testing.B) {
	for n := 0; n < b.N; n++ {
		s.Sort(data.getCopy())
	}
}

func BenchmarkInternal100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortg.Ints(intSlice100.getCopy())
	}
}

func BenchmarkInternal1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortg.Ints(intSlice1000.getCopy())
	}
}

func BenchmarkInternal10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortg.Ints(intSlice1000.getCopy())
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkSort(bs, intSlice100, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkSort(bs, intSlice1000, b)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	benchmarkSort(bs, intSlice10000, b)
}

func BenchmarkQuickSort100(b *testing.B) {
	benchmarkSort(qs, intSlice100, b)
}

func BenchmarkQuickSort1000(b *testing.B) {
	benchmarkSort(qs, intSlice1000, b)
}

func BenchmarkQuickSort10000(b *testing.B) {
	benchmarkSort(qs, intSlice10000, b)
}
