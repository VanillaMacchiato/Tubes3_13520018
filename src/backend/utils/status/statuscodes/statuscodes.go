package statuscodes

// status code
const (
	// success
	DNASuccess = "DNA_INSERTION_SUCCESS"

	// error
	InvalidDNA    = "INVALID_DNA"
	DiseaseExists = "DISEASE_ALREADY_EXISTS"
	DiseaseEmpty  = "DISEASE_EMPTY"
	ServerError   = "INTERNAL_SERVER_ERROR"
)

var text = map[string]string{
	InvalidDNA:    "DNA sequence contains illegal characters",
	DiseaseExists: "Disease name already exists",
	DiseaseEmpty:  "Disease name empty",
	ServerError:   "The server encountered an error",
}

func Text(code string) string {
	return text[code]
}
