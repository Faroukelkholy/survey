package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"faroukelkholy/survey/internal/service/survey/mocks"
)

func TestSubmitSurveyForm_Success(t *testing.T) {
	// setup mocks
	ctx := context.Background()

	sMock := &mocks.Service{}
	sMock.On("SubmitSurveyForm", ctx, mock.AnythingOfType("*models.SurveyForm")).Return(nil)

	// setup handler
	h := PSurveyFormHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/survey_forms/%s/surveys", mocks.SfID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			data := res.Data

			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "created", data.(string))
		}
	}
}

func TestSubmitSurveyForm_Error(t *testing.T) {
	// setup mocks
	ctx := context.Background()

	sMock := &mocks.Service{}
	sMock.On("SubmitSurveyForm", ctx, mock.AnythingOfType("*models.SurveyForm")).Return(errors.New("mongo: no collection defined"))
	// setup handler
	h := PSurveyFormHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/survey_forms/%s/surveys", mocks.SfID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "internal error", HTTPErr.Title)
		}
	}
}
