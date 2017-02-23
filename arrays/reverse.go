package arrays

// Reverse reverses an int slice.
func Reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
