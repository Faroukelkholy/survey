package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func seedDB(ctx context.Context, db *mongo.Database) error {
	one, err := findOneSurveyForm(ctx, db)
	if err != nil {
		return err
	}

	if one != nil {
		return nil
	}

	sf := &SurveyForm{
		ID:    primitive.NewObjectID(),
		Title: "favourite color",
		Content: []*Content{{
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

	res, err := db.Collection("survey_form").InsertOne(ctx, sf)

	if err != nil {
		log.Println("error in seedDB:", err)
		return err
	}

	s := &Survey{
		ID:           primitive.NewObjectID(),
		SurveyFormID: res.InsertedID.(primitive.ObjectID),
		Answers:      []string{"blue", "light"},
	}

	_, err = db.Collection("survey").InsertOne(ctx, s)
	return err
}

func findOneSurveyForm(ctx context.Context, db *mongo.Database) (*SurveyForm, error) {
	one := &SurveyForm{}

	result := db.Collection("survey_form").FindOne(ctx, bson.D{})
	err := result.Decode(one)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return one, nil
}
