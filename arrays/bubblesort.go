package arrays

type bubblesort struct {
}

// Sort the given array.
func (s bubblesort) Sort(array []int) {
	for j := 0; j < len(array); j++ {
		for i := 0; i < len(array); i++ {
			if i >= len(array)-1 {
				continue
			}
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}
}
