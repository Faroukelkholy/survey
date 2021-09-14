package mocks

import (
	"faroukelkholy/survey/internal/storage"
)

var (
	id, _   = storage.ObjectIDFromHex("613f62eac51520f0073254f9")
	sfID, _ = storage.ObjectIDFromHex("613f62eac51520f0073254f8")

	SurveyForm = &storage.SurveyForm{
		ID:    sfID,
		Title: "favourite color",
		Content: []*storage.Content{{
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

	Survey = &storage.Survey{
		ID:           id,
		SurveyFormID: sfID,
		Answers:      []string{"blue", "light"},
	}
)
