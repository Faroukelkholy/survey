package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type surveyRepository struct {
	db *mongo.Database
}

func NewSurveyRepository(db *mongo.Database) SurveyRepository {
	return &surveyRepository{db: db}
}

func (s surveyRepository) GetSurveysBySurveyFormID(ctx context.Context, surveyFormID string) ([]*Survey, error) {
	panic("implement me")
}

func (s surveyRepository) CreateSurvey(ctx context.Context, survey *Survey) (err error) {
	panic("implement me")
}
