package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type surveyFormRepository struct {
	db *mongo.Database
}

func NewSurveyFormRepository(db *mongo.Database) SurveyFormRepository {
	return &surveyFormRepository{db: db}
}

func (repo *surveyFormRepository) CreateSurveyForm(ctx context.Context, form *SurveyForm) (err error) {
	sf := &SurveyForm{
		ID:      primitive.NewObjectID(),
		Title:   form.Title,
		Content: form.Content,
	}

	_, err = repo.db.Collection("survey_form").InsertOne(ctx, sf)
	return
}
