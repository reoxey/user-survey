package model


type QA struct {
	Question string `json:"question" bson:"question"`
	Answer bool `json:"answer,omitempty" bson:"answer,omitempty"`
}

type Survey struct {
	Id	string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Questions []string `json:"questions,omitempty" bson:"questions,omitempty"`
}

type Result struct {
	Id	string `json:"_id,omitempty" bson:"_id,omitempty"`
	SurveyId	string `json:"survey_id,omitempty" bson:"survey_id,omitempty"`
	Title string `json:"title" bson:"title"`
	FilledAt string `json:"filled_at" bson:"filled_at"`
	Qas []*QA `json:"qas,omitempty" bson:"qas,omitempty"`
}
