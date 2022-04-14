package algorithms

// Boyer-Moore string matching algorithm.
// Returns a boolean depending if a pattern exists in the given text.
// Only searchest the first occurence.
func BoyerMoore(text, pattern string) bool {
	// initiate return var
	var found bool = false
	// initiate text read anchor head (text anchor index)
	idxAnchor := len(pattern) - 1
	// initiate text read head (text index)
	idxText := len(pattern) - 1
	// initiate pattern read head (pattern index)
	idxPattern := len(pattern) - 1
	// initiate bad match map
	badMatch := make(map[byte]int)
	badMatch = BadMatch(pattern)
	//
	for idxAnchor < len(text) && !found {
		if text[idxText] != pattern[idxPattern] {
			idxAnchor = idxAnchor + badMatch[text[idxText]]
			idxText = idxAnchor
			idxPattern = len(pattern) - 1
		} else {
			idxPattern--
			idxText--
			if idxPattern < 0 {
				found = true
			}
		}
	}

	return found
}

// Returns a map of int based on character bytes of the last occurence of each character in the given sequence.
func LastOccurence(sequence string) map[byte]int {
	chars := "AGCT"
	lastOcc := make(map[byte]int)

	// default value
	for i := 0; i < len(chars); i++ {
		lastOcc[chars[i]] = -1
	}

	// real value
	for i := 0; i < len(chars); i++ {
		for j := len(sequence) - 1; j >= 0; j-- {
			if chars[i] == sequence[j] {
				lastOcc[chars[i]] = j
				break
			}
		}
	}

	return lastOcc
}

// Returns a map of int based on character of bytes of the amount of index needs to be shuffled to the right
// if there was a bad match
func BadMatch(pattern string) map[byte]int {
	chars := "AGCT"
	badMatch := make(map[byte]int)
	lastOccurence := make(map[byte]int)
	lastOccurence = LastOccurence(pattern)

	for i := 0; i < len(chars); i++ {
		if lastOccurence[chars[i]] == -1 {
			badMatch[chars[i]] = len(pattern)
		} else {
			badMatch[chars[i]] = len(pattern) - lastOccurence[chars[i]] - 1
			if badMatch[chars[i]] == 0 {
				badMatch[chars[i]] = 1
			}
		}
	}

	return badMatch
}
