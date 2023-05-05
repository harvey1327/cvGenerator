package model

type employment struct {
	History []employmentHistory
}

type employmentHistory struct {
	Employer   string
	Position   string
	Start      string
	End        string
	Summary    string
	Highlights []string
}
