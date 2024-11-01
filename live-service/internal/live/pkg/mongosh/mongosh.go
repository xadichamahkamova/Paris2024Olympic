package mongosh

import (
	"context"
	"fmt"
	config "live-service/internal/live/pkg/load"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client     mongo.Client
	Collection mongo.Collection
}

var ctx = context.Background()

func NewConnection(cfg *config.Config) (*Mongo, error) {

	uri := fmt.Sprintf("mongodb://%s:%d", cfg.MongoConfig.Host, cfg.MongoConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	mycoll := client.Database(cfg.MongoConfig.Database).Collection(cfg.MongoConfig.Collection)
	return &Mongo{
		Client:     *client,
		Collection: *mycoll,
	}, nil
}
