package mocks

import (
	"faroukelkholy/survey/internal/service/models"
)

var (
	ID   = "613f62eac51520f0073254f9"
	SfID = "613f62eac51520f0073254f8"

	SurveyForm = &models.SurveyForm{
		ID:    SfID,
		Title: "favourite color",
		Content: []*models.Content{{
			Question: "what is your favorite color ?",
			Answers: []string{"blue",
				"green",
				"red",
				"yellow"},
		}, {
			Question: "which shade you prefer ?",
			Answers: []string{"light",
				"dark"},
		}},
	}

	Survey = &models.Survey{
		ID:           ID,
		SurveyFormID: SfID,
		Answers:      []string{"blue", "light"},
	}
)
