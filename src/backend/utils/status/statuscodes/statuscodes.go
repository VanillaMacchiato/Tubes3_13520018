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
	WrongSearchQuery = "WRONG_SEARCH_QUERY"
)

var text = map[string]string{
	InvalidDNA:       "Sekuens DNA mengandung karakter yang tidak valid",
	InvalidAlgorithm: "Algoritma yang dimasukkan tidak valid. Algoritma yang tersedia: KMP/BoyerMoore",
	DiseaseExists:    "Nama penyakit sudah ada ada pada database",
	DiseaseNotExits:  "Informasi DNA dari penyakit ini belum tersedia pada database",
	PatientNameEmpty: "Input nama pasien kosong",
	DiseaseEmpty:     "Input nama penyakit kosong",
	FileError:        "Kesalahan saat membaca file",
	ServerError:      "Terjadi kesalahan pada server",
	WrongSearchQuery: "Query yang dimasukkan salah",
}

func Text(code string) string {
	return text[code]
}
