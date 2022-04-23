package algorithms

func Damerau_Levenshtein_distance(text, pattern string) int {
	var cost, db, l, k, idxK, idxL int
	da := make(map[byte]int)
	// cuman ada {A,C,G,T}
	chars := "AGCT"
	for i := 0; i < len(chars); i++ {
		da[chars[i]] = 0
	}

	d := make([][]int, len(text)+2)
	for i := 0; i < len(text)+2; i++ {
		d[i] = make([]int, len(pattern)+2)
	}
	maxDist := len(text) + len(pattern)
	d[len(text)-1][len(pattern)-1] = maxDist
	for i := 0; i <= len(text); i++ {
		d[i][len(pattern)-1] = maxDist
		d[i][0] = i
	}
	for i := 0; i <= len(pattern); i++ {
		d[len(text)-1][i] = maxDist
		d[0][i] = i
	}

	for i := 1; i <= len(text); i++ {
		db = 0
		for j := 1; j <= len(pattern); j++ {
			k = da[pattern[j-1]]
			l = db
			if pattern[j-1] == text[i-1] {
				cost = 0
				db = j
			} else {
				cost = 1
			}
			temp := d[i-1][j-1] + cost // substitution
			if d[i][j-1]+1 < temp {    // insertion
				temp = d[i][j-1] + 1
			}
			if d[i-1][j]+1 < temp { // deletion
				temp = d[i-1][j] + 1
			}
			if k-1 < 0 {
				idxK = len(text) - (k - 1)
			} else {
				idxK = (k - 1)
			}
			if l-1 < 0 {
				idxL = len(pattern) - (l - 1)
			} else {
				idxL = (l - 1)
			}
			if d[idxK][idxL]+(i-k-1)+1+(j-l-1) < temp { // transposition
				temp = d[idxK][idxL] + (i - k - 1) + 1 + (j - l - 1)
			}
			d[i][j] = temp
		}
		da[text[i-1]] = i
	}

	return d[len(text)][len(pattern)]
}

func Likeness(text, pattern string) float32 {
	var likeness float32 = 0.0

	// check for likeness with substitution & transposition only (length of subText is the same as pattern)
	subTextLen := len(pattern)
	for i := 0; i <= len(text)-subTextLen; i++ {
		subText := text[i : i+subTextLen]
		dist := Damerau_Levenshtein_distance(subText, pattern)
		temp := float32(dist) / float32(len(pattern))
		temp = 1 - temp
		if temp < 0 {
			temp = 0.0
		}
		if likeness < temp {
			likeness = temp
		}
	}

	// check for likeness with subs, ins, del, and transposition (length of subText is 10% bigger than pattern)
	if likeness < 0.80 {
		subTextLen = int(float32(1.1) * float32(len(pattern)))
		if subTextLen > len(text) {
			subTextLen = len(text)
		}
		for i := 0; i <= len(text)-subTextLen; i++ {
			subText := text[i : i+subTextLen]
			dist := Damerau_Levenshtein_distance(subText, pattern)
			temp := float32(dist) / float32(len(pattern))
			temp = 1 - temp
			if temp < 0 {
				temp = 0.0
			}
			if likeness < temp {
				likeness = temp
			}
		}
	}
	return likeness * 100
}
