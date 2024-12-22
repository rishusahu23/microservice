package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rishu/microservice/external/enums"
	"github.com/rishu/microservice/external/post"
	"github.com/rishu/microservice/external/post/json_placeholder"
	"github.com/rishu/microservice/gen/api/rpc"
	userPb "github.com/rishu/microservice/gen/api/user"
	customerrors "github.com/rishu/microservice/pkg/errors"
	txn "github.com/rishu/microservice/pkg/transaction"
	"github.com/rishu/microservice/user/dao"
	"github.com/rishu/microservice/user/dao/mongo"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	dao        dao.UserDao
	txnManager txn.TransactionManager
	userPb.UnimplementedUserServiceServer
	postClient post.Client
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
	err := s.txnManager.RunInTxn(ctx, func(sessCtx mongo2.SessionContext) error {
		_, err := s.dao.Get(sessCtx, mongo.WithUserId(req.GetUser().GetId()))
		if err != nil && !errors.Is(err, customerrors.ErrRecordNotFound) {
			return err
		}

		// Step 3: Create user
		return s.dao.Create(sessCtx, req.GetUser())
	})

	if err != nil {
		return &userPb.CreateUserResponse{
			Status: rpc.StatusInternal(""),
		}, nil
	}
	return &userPb.CreateUserResponse{
		Status: rpc.StatusOk(),
	}, nil
}

func (s *Service) GetPost(ctx context.Context, req *userPb.GetPostRequest) (*userPb.GetPostResponse, error) {
	resp, err := s.postClient.FetchPost(ctx, &placeholder.FetchPostRequest{
		PostId: "1",
		Vendor: enums.JsonPlaceholder,
	})
	if err != nil {
		return &userPb.GetPostResponse{
			Status: rpc.StatusInternal(err.Error()),
		}, nil
	}
	return &userPb.GetPostResponse{
		Status: rpc.StatusOk(),
		Post: &userPb.Post{
			UserId: int32(resp.UserID),
			Id:     int32(resp.ID),
			Title:  resp.Title,
			Body:   resp.Body,
		},
	}, nil
}

func NewService(dao dao.UserDao, txnManager txn.TransactionManager, postClient post.Client) *Service {
	return &Service{
		dao:        dao,
		txnManager: txnManager,
		postClient: postClient,
	}
}

var _ userPb.UserServiceServer = (*Service)(nil)
