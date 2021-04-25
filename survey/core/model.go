package core

import "go.mongodb.org/mongo-driver/bson/primitive"

// QA definitive answers for each question
type QA struct {
	Question string `json:"question" bson:"question"`
	Answer bool `json:"answer" bson:"answer"`
}

// Survey stores the questions for each survey
type Survey struct {
	Id	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Questions []string `json:"questions,omitempty" bson:"questions,omitempty"`
}

// Result stores answers of the questions asked in the Survey
type Result struct {
	Id	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SurveyId	primitive.ObjectID `json:"survey_id" bson:"survey_id"`
	Title string `json:"title" bson:"title"`
	FilledAt string `json:"filled_at" bson:"filled_at"`
	Qas []*QA `json:"qas,omitempty" bson:"qas,omitempty"`
}

type Surveys []*Survey
type Results []*Result
