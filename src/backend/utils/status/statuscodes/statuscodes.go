package statuscodes

// status code
const (
	// success
	DNASuccess     = "DNA_INSERTION_SUCCESS"
	PredictSuccess = "DISEASE_PREDICTION_SUCCESS"
	ResultSuccess  = "GET_RESULT_SUCCESS"

	// error
	InvalidDNA       = "INVALID_DNA"
	InvalidAlgorithm = "INVALID_ALGORITHM"
	DiseaseExists    = "DISEASE_ALREADY_EXISTS"
	DiseaseNotExits  = "DISEASE_DOES_NOT_EXIST"
	DiseaseEmpty     = "DISEASE_EMPTY"
	PatientNameEmpty = "PATIENT_NAME_EMPTY"
	FileError        = "FILE_ERROR"
	ServerError      = "INTERNAL_SERVER_ERROR"
)

var text = map[string]string{
	InvalidDNA:       "DNA sequence contains illegal characters",
	InvalidAlgorithm: "The specified algorithm invalid. Algorithms available: KMP, BoyerMoore, or Levenshtein. ",
	DiseaseExists:    "Disease name already exists",
	DiseaseNotExits:  "DNA information from the disease requested is not yet available",
	PatientNameEmpty: "Patient name empty",
	DiseaseEmpty:     "Disease name empty",
	FileError:        "An error occured while reading the file",
	ServerError:      "The server encountered an error",
}

func Text(code string) string {
	return text[code]
}
