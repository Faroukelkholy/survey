package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type surveyRepository struct {
	db *mongo.Database
}

func NewSurveyRepository(db *mongo.Database) SurveyRepository {
	return &surveyRepository{db: db}
}

func (repo *surveyRepository) GetSurveysBySurveyFormID(ctx context.Context, surveyFormID primitive.ObjectID) (surveys []*Survey, err error) {
	filter := bson.D{primitive.E{Key: "survey_form_id", Value: surveyFormID}}

	cur, err := repo.db.Collection("survey").Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		s := &Survey{}

		err = cur.Decode(s)
		if err != nil {
			return
		}

		surveys = append(surveys, s)
	}

	err = cur.Err()
	return
}

func (repo *surveyRepository) CreateSurvey(ctx context.Context, survey *Survey) (err error) {
	s := &Survey{
		ID:           primitive.NewObjectID(),
		SurveyFormID: survey.SurveyFormID,
		Answers:      survey.Answers,
	}

	_, err = repo.db.Collection("survey").InsertOne(ctx, s)
	return
}
