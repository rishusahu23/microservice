package strategy

import (
	"context"
	"github.com/google/wire"
	"github.com/rishu/microservice/user/dao"
	"github.com/rishu/microservice/user/dao/mongo"
)

type DB struct {
	dao dao.UserDao
}

func NewDB(dao dao.UserDao) *DB {
	return &DB{dao: dao}
}

var (
	GetUserDBWireSet = wire.NewSet(NewDB, wire.Bind(new(GetUserStrategy), new(*DB)))
)

func (d *DB) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	user, err := d.dao.Get(ctx, mongo.WithUserId(req.UserId))
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{
		User: user,
	}, nil
}
