package survey

import (
	"context"

	"faroukelkholy/survey/internal/service/models"
	"faroukelkholy/survey/internal/storage"
)

//Service define services related to survey domain
type Service interface {
	SubmitSurvey(ctx context.Context, sv *models.Survey) error
	GetSurveysBySurveyFormID(ctx context.Context, surveyFormID string) ([]*models.Survey, error)
	SubmitSurveyForm(ctx context.Context, sf *models.SurveyForm) error
}

//service struct implement the Service interface
type service struct {
	repo storage.Repository
}

func New(repo storage.Repository) Service {
	return &service{repo: repo}
}

func (s *service) SubmitSurvey(ctx context.Context, sv *models.Survey) error {
	survey, err := serializeSurveyToStorage(sv)
	if err != nil {
		return err
	}

	return s.repo.CreateSurvey(ctx, survey)
}

func (s *service) GetSurveysBySurveyFormID(ctx context.Context, surveyFormID string) (svs []*models.Survey, err error) {
	sfID, err := storage.ObjectIDFromHex(surveyFormID)
	if err != nil {
		return nil, err
	}

	surveys, err := s.repo.GetSurveysBySurveyFormID(ctx, sfID)
	if err != nil {
		return nil, err
	}

	if surveys == nil {
		return
	}

	for _, sv := range surveys {
		svs = append(svs, serializeSurveyToService(sv))
	}

	return
}

func (s *service) SubmitSurveyForm(ctx context.Context, sf *models.SurveyForm) error {
	surveyForm, err := serializeSurveyFormToStorage(sf)
	if err != nil {
		return err
	}

	return s.repo.CreateSurveyForm(ctx, surveyForm)
}

// serializeSurveyToStorage translate Survey data structure from the service to the storage
func serializeSurveyToStorage(sv *models.Survey) (*storage.Survey, error) {
	sfID, err := storage.ObjectIDFromHex(sv.SurveyFormID)
	if err != nil {
		return nil, err
	}

	return &storage.Survey{
		SurveyFormID: sfID,
		Answers:      sv.Answers,
	}, nil
}

// serializeSurveyToService translate Survey data structure from the storage to the service
func serializeSurveyToService(sv *storage.Survey) *models.Survey {
	return &models.Survey{
		ID:           sv.ID.Hex(),
		SurveyFormID: sv.SurveyFormID.Hex(),
		Answers:      sv.Answers,
	}
}

// serializeSurveyFormToStorage translate SurveyFrom data structure from the service to the storage
func serializeSurveyFormToStorage(sf *models.SurveyForm) (*storage.SurveyForm, error) {
	cs := make([]*storage.Content, 0, len(sf.Content))

	for _, c := range sf.Content {
		cs = append(cs, &storage.Content{
			Question: c.Question,
			Answers:  c.Answers,
		})
	}

	return &storage.SurveyForm{
		Title:   sf.Title,
		Content: cs,
	}, nil
}
