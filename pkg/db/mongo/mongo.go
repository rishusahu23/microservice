package mongo

import (
	"context"
	"github.com/rishu/microservice/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func GetMongoClient(ctx context.Context, conf *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	readPref := readpref.Primary()

	clientOptions := options.Client().
		ApplyURI(conf.MongoConfig.MongoDBURI).
		SetReadPreference(readPref)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}
