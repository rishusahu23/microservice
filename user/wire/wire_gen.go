// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/external/post"
	redis2 "github.com/rishu/microservice/pkg/in_memory_store/redis"
	mongo3 "github.com/rishu/microservice/pkg/transaction/mongo"
	"github.com/rishu/microservice/user"
	mongo2 "github.com/rishu/microservice/user/dao/mongo"
	"github.com/rishu/microservice/user/getuserstrategy"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// Injectors from wire.go:

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client, redisClient *redis.Client) *user.Service {
	userDaoMongo := mongo2.NewUserDaoMongo(mongoClient, conf)
	mongoTransactionManager := mongo3.NewMongoTransactionManager(mongoClient)
	db := strategy.NewDB(userDaoMongo)
	redisInMemoryStore := redis2.NewRedisInMemoryStore(redisClient)
	cache := strategy.NewCache(redisInMemoryStore)
	getUserStrategyFactoryImpl := strategy.NewGetUserStrategyFactoryImpl(db, cache)
	service := user.NewService(userDaoMongo, mongoTransactionManager, getUserStrategyFactoryImpl)
	return service
}

// wire.go:

func GetPostClientProvider(provider *post.ClientImpl) post.Client {
	return provider
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
