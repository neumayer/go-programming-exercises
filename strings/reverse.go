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

func reverseInline(value []rune) {
	reverseString(value)

	var wordBeginIndex = 0
	var wordEndIndex = 0
	for readIndex := 0; readIndex < len(value); readIndex++ {
		if value[readIndex] == ' ' {
			wordEndIndex = readIndex - 1
			reverseSubString(value, wordBeginIndex, wordEndIndex)
			wordBeginIndex = readIndex + 1
		}
		if readIndex == len(value)-1 {
			wordEndIndex = readIndex
			reverseSubString(value, wordBeginIndex, wordEndIndex)
			wordBeginIndex = readIndex
		}
	}
}

func reverseString(value []rune) {
	reverseSubString(value, 0, len(value)-1)
}

func reverseSubString(value []rune, begin int, end int) {
	for ; begin < end; begin, end = begin+1, end-1 {
		value[begin], value[end] = value[end], value[begin]
	}
}
