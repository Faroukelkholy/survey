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

	"faroukelkholy/survey/internal/service/models"
	"faroukelkholy/survey/internal/service/survey/mocks"
)

func TestGSurveysBySfIDHandler_Success(t *testing.T) {
	// setup mocks
	ctx := context.Background()

	surveys := []*models.Survey{mocks.Survey}

	sMock := &mocks.Service{}
	sMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("string")).Return(surveys, nil)

	// setup handler
	h := GSurveysBySfIDHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/survey_forms/%s/surveys", mocks.SfID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			data := res.Data

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, surveys[0].ID, data.([]interface{})[0].(map[string]interface{})["id"])
			assert.Equal(t, surveys[0].SurveyFormID, data.([]interface{})[0].(map[string]interface{})["survey_form_id"])
		}
	}
}

func TestGSurveysBySfIDHandler_Error(t *testing.T) {
	// setup mocks
	ctx := context.Background()

	sMock := &mocks.Service{}
	sMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("string")).Return(nil, errors.New("mongo: no collection defined"))

	// setup handler
	h := GSurveysBySfIDHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/survey_forms/%s/surveys", mocks.SfID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "internal error", HTTPErr.Title)
			assert.Equal(t, "", HTTPErr.Description)
		}
	}
}

func TestGSurveysBySfIDHandler_NotFound(t *testing.T) {
	// setup mocks
	ctx := context.Background()

	sMock := &mocks.Service{}
	sMock.On("GetSurveysBySurveyFormID", ctx, mock.AnythingOfType("string")).Return(nil, nil)

	// setup handler
	h := GSurveysBySfIDHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/survey_forms/%s/surveys", mocks.SfID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "surveys not found", HTTPErr.Title)
			assert.Equal(t, "", HTTPErr.Description)
		}
	}
}
