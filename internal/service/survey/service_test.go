package survey

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	modelMock "faroukelkholy/survey/internal/service/survey/mocks"
	"faroukelkholy/survey/internal/storage"
	"faroukelkholy/survey/internal/storage/mocks"
)

func TestGetSurveysBySurveyFormID_Success(t *testing.T) {
	ctx := context.Background()
	rMock := &mocks.Repository{}
	rMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("primitive.ObjectID")).Return([]*storage.Survey{mocks.Survey}, nil)

	s := New(rMock)
	result, err := s.GetSurveysBySurveyFormID(ctx, modelMock.SfID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, mocks.Survey.ID.Hex(), result[0].ID)
	assert.EqualValues(t, mocks.Survey.SurveyFormID.Hex(), result[0].SurveyFormID)
}

func TestGetSurveysBySurveyFormID_Error(t *testing.T) {
	ctx := context.Background()

	var cases = []struct {
		title string
	}{
		{
			"caseNoRows",
		},
		{
			"caseDBError",
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			switch test.title {
			case "caseNoRows":
				rMock := &mocks.Repository{}
				rMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("primitive.ObjectID")).Return(nil, nil)

				s := New(rMock)
				result, err := s.GetSurveysBySurveyFormID(ctx, modelMock.SfID)

				assert.Nil(t, err)
				assert.Nil(t, result)
			case "caseDBError":
				rMock := &mocks.Repository{}
				rMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("primitive.ObjectID")).Return(nil, errors.New("mongo: no documents in result"))

				s := New(rMock)
				result, err := s.GetSurveysBySurveyFormID(ctx, modelMock.SfID)

				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSubmitSurvey_Success(t *testing.T) {
	ctx := context.Background()

	rMock := &mocks.Repository{}
	rMock.On("CreateSurvey", ctx, mock.AnythingOfType("*storage.Survey")).Return(nil)

	s := New(rMock)
	err := s.SubmitSurvey(ctx, modelMock.Survey)

	assert.Nil(t, err)
}

func TestSubmitSurvey_Error(t *testing.T) {
	ctx := context.Background()

	rMock := &mocks.Repository{}
	rMock.On("CreateSurvey", ctx, mock.AnythingOfType("*storage.Survey")).Return(errors.New("mongo: no collection defined"))

	s := New(rMock)
	err := s.SubmitSurvey(ctx, modelMock.Survey)

	assert.NotNil(t, err)
}

func TestSubmitSurveyForm_Success(t *testing.T) {
	ctx := context.Background()

	rMock := &mocks.Repository{}
	rMock.On("CreateSurveyForm", ctx, mock.AnythingOfType("*storage.SurveyForm")).Return(nil)

	s := New(rMock)
	err := s.SubmitSurveyForm(ctx, modelMock.SurveyForm)

	assert.Nil(t, err)
}

func TestSubmitSurveyForm_Error(t *testing.T) {
	ctx := context.Background()

	rMock := &mocks.Repository{}
	rMock.On("CreateSurveyForm", ctx, mock.AnythingOfType("*storage.SurveyForm")).Return(errors.New("mongo: no collection defined"))

	s := New(rMock)
	err := s.SubmitSurveyForm(ctx, modelMock.SurveyForm)

	assert.NotNil(t, err)
}
