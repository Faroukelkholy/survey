package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SurveyForm struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Content []*Content         `bson:"content"`
}

type Survey struct {
	ID           primitive.ObjectID `bson:"_id"`
	SurveyFormID primitive.ObjectID `bson:"survey_form_id"`
	Answers      []string           `bson:"answers"`
}

type Content struct {
	Question string   `bson:"question"`
	Answers  []string `bson:"answers"`
}

func HexFromString(s string) string {
	var bytes [12]byte
	copy(bytes[:], s)

	return primitive.ObjectID(bytes).Hex()
}

func ObjectIDFromHex(hex string) (objID primitive.ObjectID, err error) {
	objID, err = primitive.ObjectIDFromHex(hex)
	return
}
