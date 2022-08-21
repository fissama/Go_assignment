package repositories

import (
	"Go/configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"time"
)

type (
	Assignment struct {
		Number   int32     `json:"number"`
		InsertAt time.Time `json:"insertAt"`
	}
)

func InsertAssignment(number string, client *mongo.Client) bool {
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	newAssignment := &Assignment{
		Number:   int32(n),
		InsertAt: time.Now().UTC(),
	}

	config := configs.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	numbersCollection := client.Database(config.DB).Collection(config.Collection)

	numbersCollection.InsertOne(ctx, newAssignment)

	return true
}
