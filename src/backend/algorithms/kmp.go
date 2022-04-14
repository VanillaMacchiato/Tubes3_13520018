package algorithms

import "fmt"

func KMP(text, pattern string) []int {
	fmt.Printf("text: %v\n", text)
	fmt.Printf("pattern: %v\n", pattern)

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
			fmt.Printf("Found at index %d\n", i-j)
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

	fmt.Println(foundList)
	return foundList
}

