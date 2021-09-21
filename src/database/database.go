package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(mongoURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connects to MongoDB
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// ping to ensure connectivity
	err = client.Ping(ctx, readpref.Nearest(readpref.WithMaxStaleness(90*time.Second)))
	if err != nil {
		return nil, err
	}

	return client, nil
}
