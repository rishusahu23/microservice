package mongo

import (
	"context"
	"flag"
	"github.com/rishu/microservice/config"
	pkgMongo "github.com/rishu/microservice/pkg/test/mongo"
	mongo2 "github.com/rishu/microservice/user/dao/models/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	ctx := context.Background()
	conf, mongoClient, teardown := Initialise(ctx)
	defer teardown()
	UserDaoMongoTS = NewUserDaoMongo(mongoClient, conf)

	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func Initialise(ctx context.Context) (*config.Config, *mongo.Client, func()) {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongoClient := pkgMongo.SetUpTestDB(conf, mongo2.UserCollectionName)
	return conf, mongoClient, func() {
		_ = mongoClient.Disconnect(ctx)
	}
}
