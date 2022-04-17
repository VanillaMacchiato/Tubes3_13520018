package algorithms

func KMP(text, pattern string) []int {

	// border function
	patternLength := len(pattern)
	b := make([]int, patternLength)
	b[0] = 0
	for i := 1; i < patternLength; i++ {
		maxLen := 0
		for j := 1; j < i+1; j++ { // length of sequence
			if pattern[0:j] == pattern[(i-j+1):i+1] {
				maxLen = j
			}
		}
		b[i] = maxLen
	}

	// text-pattern matching
	i := 0
	j := 0
	textLength := len(text)
	var foundList []int
	for i < textLength {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == patternLength {
			foundList = append(foundList, i-j)
			j = b[j-1]
		} else if i < textLength && text[i] != pattern[j] {
			if j > 0 {
				j = b[j-1]
			} else {
				i++
			}
		}
	}

	return foundList
}
