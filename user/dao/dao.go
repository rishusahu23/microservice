package dao

import (
	"context"
	"github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/pkg/filters"
)

// mockgen -source=user/dao/dao.go -destination=user/mocks/dao/dao.go -package=mocks

type UserDao interface {
	Get(ctx context.Context, filters ...filters.FilterOption) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}
