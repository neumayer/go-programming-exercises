package arrays

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var quicksortTests = []struct {
		n        []int
		expected []int
	}{
		{[]int{6, 5, 4, 1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 1, 2, 3, 2}, []int{1, 2, 2, 3, 4, 5, 6}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}
	for _, tt := range quicksortTests {
		original := make([]int, len(tt.n))
		copy(original, tt.n)
		bs := quicksort{}
		bs.Sort(tt.n)

		if !reflect.DeepEqual(tt.n, tt.expected) {
			t.Errorf("Sort(%d): expected %d, actual %d", original, tt.expected, tt.n)
		}
	}
}
