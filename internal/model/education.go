package model

type education struct {
	History []educationHistory
}

type educationHistory struct {
	Institution string
	Area        string
	StudyType   string
	Start       string
	End         string
}
