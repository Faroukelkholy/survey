package models

type Survey struct {
	ID           string   `json:"id,omitempty"`
	SurveyFormID string   `json:"survey_form_id"`
	Answers      []string `json:"answers"`
}
