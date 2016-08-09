package arrays

import "github.com/neumayer/go-programming-exercises/utils"

type bubblesort struct {
}

// Sort the given array.
func (s bubblesort) Sort(array []int) {
	for i := 0; i < len(array); i++ {
		swapPerformed := false
		for i := 0; i < len(array); i++ {
			if i+1 == len(array) {
				break
			}
			if array[i] > array[i+1] {
				arrayutils.Swap(array, i, i+1)
				swapPerformed = true
			}
		}
		if !swapPerformed {
			break
		}
	}
}
