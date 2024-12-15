package dao

import (
	"context"
	"github.com/google/wire"
	"github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/user/dao/mongo"
)

var (
	UserDaoWireSet = wire.NewSet(mongo.NewUserDaoMongo, wire.Bind(new(UserDao), new(*mongo.UserDaoMongo)))
)

type UserDao interface {
	Get(ctx context.Context, userId string) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}
