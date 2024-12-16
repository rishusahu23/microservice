package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rishu/microservice/gen/api/rpc"
	userPb "github.com/rishu/microservice/gen/api/user"
	customerrors "github.com/rishu/microservice/pkg/errors"
	"github.com/rishu/microservice/user/dao"
	"github.com/rishu/microservice/user/dao/mongo"
)

type Service struct {
	dao dao.UserDao
	userPb.UnimplementedUserServiceServer
}

func (s *Service) GetUser(ctx context.Context, req *userPb.GetUserRequest) (*userPb.GetUserResponse, error) {
	user, err := s.dao.Get(ctx, mongo.WithUserId(req.GetUserId()))
	fmt.Println(user)
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			return &userPb.GetUserResponse{
				Status: rpc.StatusRecordNotFound(""),
			}, nil
		}
		return &userPb.GetUserResponse{
			Status: rpc.StatusInternal(""),
		}, nil
	}
	return &userPb.GetUserResponse{
		Status: rpc.StatusOk(),
		User:   user,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error) {
	req.User.Id = uuid.NewString()
	if err := s.dao.Create(ctx, req.User); err != nil {
		return &userPb.CreateUserResponse{
			Status: rpc.StatusInternal(""),
		}, nil
	}
	return &userPb.CreateUserResponse{
		Status: rpc.StatusOk(),
	}, nil
}

func NewService(dao dao.UserDao) *Service {
	return &Service{
		dao: dao,
	}
}

var _ userPb.UserServiceServer = (*Service)(nil)
