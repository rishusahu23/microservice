package dao

import "context"

type User interface {
	CreateUser(ctx context.Context)
}
