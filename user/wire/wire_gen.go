// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/rishu/microservice/config"
	mongo3 "github.com/rishu/microservice/pkg/transaction/mongo"
	"github.com/rishu/microservice/user"
	mongo2 "github.com/rishu/microservice/user/dao/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client) *user.Service {
	userDaoMongo := mongo2.NewUserDaoMongo(mongoClient, conf)
	mongoTransactionManager := mongo3.NewMongoTransactionManager(mongoClient)
	service := user.NewService(userDaoMongo, mongoTransactionManager)
	return service
}
