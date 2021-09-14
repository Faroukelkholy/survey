package models

type SurveyForm struct {
	ID      string     `json:"id,omitempty"`
	Title   string     `json:"title"`
	Content []*Content `json:"content"`
}

type Content struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}
