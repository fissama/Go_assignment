package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type (
	Assignment struct {
		Number   int32     `json:"number"`
		InsertAt time.Time `json:"insertAt"`
	}
)

func InsertAssignment(ctx context.Context, number string, numbersCollection *mongo.Collection) (bool, error) {
	n, err := strconv.Atoi(number)
	if err != nil {
		return false, err
	}

	newAssignment := &Assignment{
		Number:   int32(n),
		InsertAt: time.Now().UTC(),
	}

	numbersCollection.InsertOne(ctx, newAssignment)

	return true, nil
}
