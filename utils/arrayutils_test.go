package arrayutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPositiveCases(t *testing.T) {
	var reverseTests = []struct {
		n        []int
		i        int
		j        int
		expected []int
	}{
		{[]int{1, 2, 3}, 1, 2, []int{1, 3, 2}},
		{[]int{1, 2, 3}, 0, 1, []int{2, 1, 3}},
		{[]int{1, 2, 3}, 0, 2, []int{3, 2, 1}},
	}

	for _, tt := range reverseTests {
		err := Swap(tt.n, tt.i, tt.j)

		if err != nil || !reflect.DeepEqual(tt.n, tt.expected) {
			t.Errorf("Swap(%d, %d, %d): expected %d, actual %d", tt.n, tt.i, tt.j, tt.expected, tt.n)
		}
	}
}

func TestNegativeCases(t *testing.T) {
	var reverseTests = []struct {
		n []int
		i int
		j int
	}{
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 2, 3}, 0, 3},
		{[]int{1, 2, 3}, -1, 3},
	}

	for _, tt := range reverseTests {
		fmt.Println(tt)
		err := Swap(tt.n, tt.i, tt.j)

		if err == nil {
			t.Errorf("Swap(%d, %d, %d): expected %d, actual %d", tt.n, tt.i, tt.j, tt.n)
		}
	}
}
