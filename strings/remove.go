package strings

// getLookup converts a string into an asci lookup table.
func getLookup(value string) []bool {
	lookup := make([]bool, 128)
	for _, v := range value {
		lookup[v] = true
	}
	return lookup
}

func remove(value []rune, remove string) string {
	lookup := getLookup(remove)

	var writeIndex int
	for _, v := range value {
		if !lookup[v] {
			value[writeIndex] = v
			writeIndex++
		}
	}
	return string(value[:writeIndex])
}
