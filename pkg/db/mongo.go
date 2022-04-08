package db

import (
	"context"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	*mongo.Collection
}

func NewMongoDB(logger LoggerInterface) *mongoDB {
	for retries := 10; retries > 0; retries-- {
		client, err := mongo.Connect(context.Background(), options.Client().
			ApplyURI("mongodb://root:really_safe_pw@localhost:27017"))
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		return &mongoDB{
			client.Database("local").Collection("trivia"),
		}
	}
	logger.Fatal("Database connection could not be established.")
	return &mongoDB{} // this should never happen since Fatal is supposed to terminate the app
}

func (m *mongoDB) Find(filters map[string]string) (results []Entry, err error) {
	filterDocument := bson.M{}
	if value, exists := filters["text"]; exists {
		values := strings.Split(value, ",")
		list := bson.A{}
		for _, singleValue := range values {
			list = append(list, primitive.Regex{Pattern: singleValue})
		}
		filterDocument["text"] = bson.D{
			{Key: "$in", Value: list},
		}
	}
	if value, exists := filters["number"]; exists {
		values := strings.Split(value, ",")
		list := bson.A{}
		for _, singleValue := range values {
			num, _ := strconv.ParseFloat(singleValue, 64)
			list = append(list, num)
		}
		filterDocument["number"] = bson.D{
			{Key: "$in", Value: list},
		}
	}
	cursor, err := m.Collection.Find(context.Background(), filterDocument)
	if err == nil {
		cursor.All(context.Background(), &results)
	}
	return results, err
}
