package arrays

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var bubblesortTests = []struct {
		n        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}
	for _, tt := range bubblesortTests {
		original := make([]int, len(tt.n))
		copy(original, tt.n)
		bs := bubblesort{}
		bs.Sort(tt.n)

		if !reflect.DeepEqual(tt.n, tt.expected) {
			t.Errorf("Sort(%d): expected %d, actual %d", original, tt.expected, tt.n)
		}
	}
}
