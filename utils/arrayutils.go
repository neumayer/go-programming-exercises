package arrayutils

import "fmt"

// Swap swaps elements in the given array.
// Returns error if either one is out of range, the lower index is higher than the higher index,
// or if they are identical.
func Swap(array []int, lowerIndex int, upperIndex int) (err error) {
	length := len(array)
	if lowerIndex >= upperIndex || upperIndex > length-1 || lowerIndex < 0 {
		return fmt.Errorf("")
	}
	iVal := array[lowerIndex]
	jVal := array[upperIndex]
	array[lowerIndex] = jVal
	array[upperIndex] = iVal
	return nil
}
