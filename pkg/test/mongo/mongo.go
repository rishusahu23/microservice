package mongo

import (
	"context"
	"github.com/rishu/microservice/config"
	pkgMongo "github.com/rishu/microservice/pkg/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpTestDB(conf *config.Config, collectionStr string) *mongo.Client {
	ctx := context.Background()
	client := pkgMongo.GetMongoClient(ctx, conf)
	collection := client.Database(conf.MongoConfig.MongoDBName).
		Collection(collectionStr)
	if err := collection.Drop(ctx); err != nil {
		panic(err)
	}
	return client
}
