package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	SurveyFormRepository
	SurveyRepository
}

type SurveyFormRepository interface {
	CreateSurveyForm(context.Context, *SurveyForm) (err error)
}

type SurveyRepository interface {
	GetSurveysBySurveyFormID(ctx context.Context, surveyFormID primitive.ObjectID) ([]*Survey, error)
	CreateSurvey(context.Context, *Survey) (err error)
}
