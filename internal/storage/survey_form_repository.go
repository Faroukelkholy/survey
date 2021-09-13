package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type surveyFormRepository struct {
	db *mongo.Database
}

func NewSurveyFormRepository(db *mongo.Database) SurveyFormRepository {
	return &surveyFormRepository{db: db}
}

func (s *surveyFormRepository) CreateSurveyForm(ctx context.Context, from *SurveyFrom) (err error) {
	panic("implement me")
}
