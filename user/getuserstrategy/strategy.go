package strategy

import (
	"context"
	"github.com/rishu/microservice/gen/api/user"
)

type GetUserStrategy interface {
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
}

type GetUserRequest struct {
	UserId string
}

type GetUserResponse struct {
	User *user.User
}
