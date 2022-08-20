package services

import (
	"Go/configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type (
	SampleData struct {
		Number   int       `json:"number"`
		InsertAt time.Time `json:"insertAt"`
	}
)

func InsertSampleData(n int) bool {
	newSampleData := &SampleData{
		Number:   n,
		InsertAt: time.Now().UTC(),
	}

	config := configs.GetConfig()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + config.User + ":" + config.Password + "@" + config.Cluster + "/test?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	numbers := client.Database(config.DB).Collection(config.Collection)

	numbers.InsertOne(ctx, newSampleData)

	return true
}
