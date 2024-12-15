//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/user"
	"github.com/rishu/microservice/user/dao"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client) *user.Service {
	wire.Build(
		user.NewService,
		dao.UserDaoWireSet,
	)
	return &user.Service{}
}
