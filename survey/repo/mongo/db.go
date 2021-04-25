package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"survey/core"
)

type repo struct {
	db *mongo.Database
}

func (r repo) CreateSurvey(ctx context.Context, survey *core.Survey) (string, error) {
	col := r.db.Collection("surveys")

	res, err := col.InsertOne(ctx, survey)
	if err != nil {
		return "", err
	}

	// returns the mongo id as a string
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("not inserted")
}

func (r repo) GetAllSurveys(ctx context.Context) (core.Surveys, error) {
	col := r.db.Collection("surveys")

	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var surveys core.Surveys
	for cur.Next(ctx) {
		s := &core.Survey{}
		cur.Decode(&s)
		surveys = append(surveys, s)
	}

	if err = cur.Err(); err != nil { // checks if any error occurred during the iteration
		return nil, err
	}

	return surveys, nil
}

func (r repo) GetSurvey(ctx context.Context, surveyId string) (*core.Survey, error) {
	col := r.db.Collection("surveys")

	id, err := primitive.ObjectIDFromHex(surveyId)
	if err != nil {
		return nil, err
	}
	var survey *core.Survey
	err = col.FindOne(ctx, bson.M{"_id": id}).Decode(&survey)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {  // error ignored if no documents
			return nil, nil
		}
		return nil, err
	}
	return survey, nil
}

func (r repo) AddResult(ctx context.Context, result *core.Result, surveyId string) (string, error) {
	col := r.db.Collection("results")

	id, err := primitive.ObjectIDFromHex(surveyId)
	if err != nil {
		return "", err
	}
	result.SurveyId = id

	res, err := col.InsertOne(ctx, result)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("not inserted")
}

func (r repo) GetAllResultsBySurvey(ctx context.Context, surveyId string) (core.Results, error) {
	col := r.db.Collection("results")

	id, err := primitive.ObjectIDFromHex(surveyId)
	if err != nil {
		return nil, err
	}
	cur, err := col.Find(ctx, bson.M{"survey_id": id})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results core.Results
	for cur.Next(ctx) {
		x := &core.Result{}
		cur.Decode(&x)
		results = append(results, x)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	fmt.Println("->", results)

	return results, nil
}

func (r repo) GetResultById(ctx context.Context, resultId string) (*core.Result, error) {
	col := r.db.Collection("results")

	id, err := primitive.ObjectIDFromHex(resultId)
	if err != nil {
		return nil, err
	}
	var result *core.Result
	err = col.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" { // error ignored if no documents
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// NewRepo initialised mongodb repo and connects to the mongodb server
func NewRepo(ctx context.Context, dsn, db string) (core.Repo, error) {

	var (
		client *mongo.Client
		err error
	)

	// mongodb connection attempt with sleep delay
	// context in main will timeout anyways if mongo remained stalled for a long time
	for i := 0; i < 3; i++ {
		clientOptions := options.Client().ApplyURI(dsn)
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			time.Sleep(2*time.Second)
			continue
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			time.Sleep(1*time.Second)
			continue
		}
		err = nil
		break
	}

	if err != nil {
		return nil, err
	}

	return &repo{
		client.Database(db),
	}, nil
}
