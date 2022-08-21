package mongodb

import (
	"Go/configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connectTimeout  = 10000 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 100
)

func NewMongoDBConnect(ctx context.Context) (*mongo.Client, error) {
	//Get config in ./configs/config_dev.json
	config := configs.GetConfig()

	//Init mongo db options
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + config.User + ":" + config.Password + "@" + config.Cluster + "/test?retryWrites=true&w=majority").
		SetConnectTimeout(connectTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetServerAPIOptions(serverAPIOptions)

	//Start opening mongo db connection
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
