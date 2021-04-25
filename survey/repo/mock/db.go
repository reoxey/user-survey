package mock

import (
	"context"

	"survey/core"
)

type repo struct {
}

func (r repo) CreateSurvey(ctx context.Context, survey *core.Survey) (string, error) {
	return "xxxxx", nil
}

func (r repo) GetAllSurveys(ctx context.Context) (core.Surveys, error) {
	return core.Surveys{
		{
			Title: "One",
		},
		{
			Title: "Two",
		},
	}, nil
}

func (r repo) GetSurvey(ctx context.Context, surveyId string) (*core.Survey, error) {
	return &core.Survey{
		Title:     "Music",
		CreateAt:  "2021",
		Questions: []string{"Do you like instrumental music?"},
	}, nil
}

func (r repo) AddResult(ctx context.Context, result *core.Result, surveyId string) (string, error) {
	return "xxxxx", nil
}

func (r repo) GetAllResultsBySurvey(ctx context.Context, surveyId string) (core.Results, error) {
	return core.Results{}, nil
}

func (r repo) GetResultById(ctx context.Context, resultId string) (*core.Result, error) {
	return &core.Result{}, nil
}

// NewMock initialised a mock db to pass unit tests
func NewMock() core.Repo {
	return &repo{}
}

