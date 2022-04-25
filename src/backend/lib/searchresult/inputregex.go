package searchresult

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func SanitizeInput(input string) (string, string, error) {
	trueInput := strings.Replace(input, "%20", " ", -1)

	date, disease := "", ""

	regex, err := regexp.Compile(`(\b[0-9]{4}-[0-9]{2}-[0-9]{2}\b)|(\b[0-9]{2}-[0-9]{2}-[0-9]{4}\b)`)
	if err != nil {
		fmt.Printf("Regular expression error: %s", err.Error())
		return "", "", errors.New("Regular expression error")
	}

	dateRegex := regex.FindAllString(trueInput, -1)
	if len(dateRegex) == 1 {
		date = dateRegex[0]
		trueInput = regex.ReplaceAllString(trueInput, "")
	} else {
		fmt.Printf("Invalid date format, has %d date inputs", len(dateRegex))
	}

	regex, err = regexp.Compile(`(\b[\S]+\b)`)
	if err != nil {
		fmt.Printf("Regular expression error: %s", err.Error())
		return "", "", errors.New("Regular expression error")
	}

	diseaseRegex := regex.FindAllString(trueInput, -1)
	if len(diseaseRegex) == 1 {
		disease = diseaseRegex[0]
	} else {
		fmt.Printf("Invalid disease format, has %d disease inputs", len(dateRegex))
	}

	return date, disease, nil
}
