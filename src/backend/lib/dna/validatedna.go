package dna

import (
	"fmt"
	"regexp"
)

func ValidateDNA(input string) bool {
	regex, err := regexp.Compile(`[^ACGT]`)

	if err != nil {
		fmt.Printf("Regular expression error: %s", err.Error())
		return false
	} else {
		res := regex.FindString(input)
		return res == ""
	}
}
