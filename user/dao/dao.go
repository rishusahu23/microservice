package dao

import (
	"context"
	"github.com/rishu/microservice/gen/api/user"
)

type UserDao interface {
	Get(ctx context.Context, userId string) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}
