package algorithms

import "fmt"

func KMP(text, pattern string) []int {
	fmt.Printf("text: %v\n", text)
	fmt.Printf("pattern: %v\n", pattern)

	// border function
	p_length := len(pattern)
	b := make([]int, p_length)
	b[0] = 0
	for i := 1; i < p_length; i++ {
		max_len := 0
		for j := 1; j < i+1; j++ { // length of sequence
			if pattern[0:j] == pattern[(i-j+1):i+1] {
				max_len = j
			}
		}
		b[i] = max_len
	}

	// text-pattern matching
	i := 0
	j := 0
	t_length := len(text)
	var found_list []int
	for i < t_length {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == p_length {
			fmt.Printf("Found at index %d\n", i-j)
			found_list = append(found_list, i-j)
			j = b[j-1]
		} else if i < t_length && text[i] != pattern[j] {
			if j > 0 {
				j = b[j-1]
			} else {
				i++
			}
		}
	}

	fmt.Println(found_list)
	return found_list
}

