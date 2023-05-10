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
	Short      bool
}

func (eh *employmentHistory) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type raw employmentHistory
	value := (*raw)(eh)
	if err := unmarshal(value); err != nil {
		return err
	}
	if value.End == "" {
		eh.End = "Present"
	}
	return nil
}
