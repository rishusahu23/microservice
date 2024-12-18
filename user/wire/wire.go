//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/rishu/microservice/config"
	mongo2 "github.com/rishu/microservice/pkg/transaction/mongo"
	"github.com/rishu/microservice/user"
	mongoDao "github.com/rishu/microservice/user/dao/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client) *user.Service {
	wire.Build(
		user.NewService,
		mongoDao.UserDaoWireSet,
		mongo2.DefaultTxnManagerWireSet,
	)
	return &user.Service{}
}
