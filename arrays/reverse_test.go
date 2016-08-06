package arrays

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var reverseTests = []struct {
		n        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range reverseTests {
		Reverse(tt.n)

		if !reflect.DeepEqual(tt.n, tt.expected) {
			t.Errorf("Reverse(%d): expected %d, actual %d", tt.n, tt.expected, tt.n)
		}
	}
}
