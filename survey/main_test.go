package main_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"survey/core"
	"survey/logger"
	"survey/repo/mock"
	"survey/router"
)

func ginServer() *gin.Engine {

	log := logger.New()

	service := core.NewService(mock.NewMock())

	r := router.New(log, true)

	r.Handle(service)

	return r.Engine
}

func TestSurvey(t *testing.T){

	client := http.Client{}

	ts := httptest.NewServer(ginServer())
	defer ts.Close()

	tests := []struct {
		reason string
		endpoint string
		method string
		status int
		payload io.Reader
	}{
		{
			"Should create a new survey with status 201",
			"%s/api/surveys",
			http.MethodPost,
			http.StatusCreated,
			bytes.NewBuffer([]byte(`{"title": "", "questions": []}`)),
		},
		{
			"Should fetch all the surveys with status 200",
			"%s/api/surveys",
			http.MethodGet,
			http.StatusOK,
			nil,
		},
		{
			"Should fetch one survey by id with status 200",
			"%s/api/surveys/qwerty",
			http.MethodGet,
			http.StatusOK,
			nil,
		},
	}

	for _, test := range tests {

		t.Run(test.reason, func(t *testing.T) {

			req, err := http.NewRequest(test.method, fmt.Sprintf(test.endpoint, ts.URL), test.payload)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)

			} else if resp.StatusCode != test.status {
				t.Errorf("Expected status code %d, got %v", test.status, resp.StatusCode)
			}
			defer resp.Body.Close()
		})
	}
}

func TestResult(t *testing.T){

	client := http.Client{}

	ts := httptest.NewServer(ginServer())
	defer ts.Close()

	tests := []struct {
		reason string
		endpoint string
		method string
		status int
		payload io.Reader
	}{
		{
			"Should add a result to a survey with status 201",
			"%s/api/surveys/qwerty/results",
			http.MethodPost,
			http.StatusCreated,
			bytes.NewBuffer([]byte(`{"title": "", "questions": []}`)),
		},
		{
			"Should fetch all the results for a survey with status 200",
			"%s/api/surveys/qwerty/results",
			http.MethodGet,
			http.StatusOK,
			nil,
		},
		{
			"Should fetch one result for a survey by result_id with status 200",
			"%s/api/surveys/qwerty/results/asdfghjkl",
			http.MethodGet,
			http.StatusOK,
			nil,
		},
	}

	for _, test := range tests {

		t.Run(test.reason, func(t *testing.T) {

			req, err := http.NewRequest(test.method, fmt.Sprintf(test.endpoint, ts.URL), test.payload)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)

			} else if resp.StatusCode != test.status {
				t.Errorf("Expected status code %d, got %v", test.status, resp.StatusCode)
			}
			defer resp.Body.Close()
		})
	}
}
