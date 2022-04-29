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
		return "", "", errors.New("Err: Regular Expression")
	}

	dateRegex := regex.FindAllString(trueInput, -1)
	if len(dateRegex) == 1 {
		regex, err = regexp.Compile(`(\b[0-9]{4}-[0-9]{2}-[0-9]{2}\b)`)
		dateRegex = regex.FindAllString(trueInput, -1)
		if len(dateRegex) == 1 {
			date = dateRegex[0]
		} else {
			regex, err = regexp.Compile(`(\b[0-9]{2}-[0-9]{2}-[0-9]{4}\b)`)
			dateRegex = regex.FindAllString(trueInput, -1)
			date = dateRegex[0]
			date = date[6:10] + "-" + date[3:5] + "-" + date[0:2]
		}
		trueInput = regex.ReplaceAllString(trueInput, "")
	} else if len(dateRegex) > 1 {
		return "", "", errors.New("Terdapat lebih dari 1 tanggal pada input")
	}

	regex, err = regexp.Compile(`(\b[\S]+\b)`)
	if err != nil {
		fmt.Printf("Regular expression error: %s", err.Error())
		return "", "", errors.New("Regular expression error")
	}

	diseaseRegex := regex.FindAllString(trueInput, -1)
	if len(diseaseRegex) == 1 {
		disease = diseaseRegex[0]
	} else if len(diseaseRegex) > 1 {
		return "", "", errors.New("Terdapat lebih dari 1 penyakit pada input")
	}

	return date, disease, nil
}
