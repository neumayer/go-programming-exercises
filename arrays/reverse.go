package arrays

import (
	"github.com/neumayer/go-programming-exercises/utils"
)

// Reverse reverses an array.
func Reverse(array []int) {
	lastIndex := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		arrayutils.Swap(array, i, lastIndex-i)
	}
}
