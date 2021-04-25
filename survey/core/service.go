package core

import "context"

// Service interface implemented by application layer
type Service interface {
	CreateSurvey(ctx context.Context, survey *Survey) (string, error)
	GetAllSurveys(ctx context.Context) (Surveys, error)
	GetSurvey(ctx context.Context, surveyId string) (*Survey, error)

	AddResult(ctx context.Context, result *Result, surveyId string) (string, error)
	GetAllResultsBySurvey(ctx context.Context, surveyId string) (Results, error)
	GetResultById(ctx context.Context, resultId string) (*Result, error)
}
