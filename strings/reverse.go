package strings

func reverse(value []rune) string {
	resultSlice := make([]rune, len(value))
	var readIndex = len(value) - 1
	var writeIndex = 0
	var wordEndIndex = 0
	var wordBeginIndex = 0
	for readIndex > 0 {
		if value[readIndex] == ' ' {
			resultSlice[writeIndex] = value[readIndex]
			writeIndex++
			readIndex--
		} else {
			wordEndIndex = readIndex + 1
			for readIndex >= 0 && value[readIndex] != ' ' {
				readIndex--
			}
			wordBeginIndex = readIndex + 1
			for wordBeginIndex < wordEndIndex {
				resultSlice[writeIndex] = value[wordBeginIndex]
				wordBeginIndex++
				writeIndex++
			}
		}
	}
	return string(resultSlice)
}
