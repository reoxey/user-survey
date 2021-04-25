package core

import (
	"context"
	"time"
)

// surveyService implements Service which accepts requests from handlers and processes the data
// repo will persist the data
type surveyService struct {
	repo Repo
}

func (s surveyService) CreateSurvey(ctx context.Context, survey *Survey) (string, error) {
	survey.CreateAt = time.Now().String()
	return s.repo.CreateSurvey(ctx, survey)
}

func (s surveyService) GetAllSurveys(ctx context.Context) (Surveys, error) {
	return s.repo.GetAllSurveys(ctx)
}

func (s surveyService) GetSurvey(ctx context.Context, surveyId string) (*Survey, error) {
	return s.repo.GetSurvey(ctx, surveyId)
}

func (s surveyService) AddResult(ctx context.Context, result *Result, surveyId string) (string, error) {
	result.FilledAt = time.Now().String()
	return s.repo.AddResult(ctx, result, surveyId)
}

func (s surveyService) GetAllResultsBySurvey(ctx context.Context, surveyId string) (Results, error) {
	return s.repo.GetAllResultsBySurvey(ctx, surveyId)
}

func (s surveyService) GetResultById(ctx context.Context, resultId string) (*Result, error) {
	return s.repo.GetResultById(ctx, resultId)
}

func NewService(repo Repo) Service {
	return &surveyService{repo}
}
