package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rishu/microservice/gen/api/rpc"
	userPb "github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/gen/api/user/enums"
	customerrors "github.com/rishu/microservice/pkg/errors"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *userPb.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		mocks   func(md *MockDependencies)
		want    *userPb.CreateUserResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &userPb.CreateUserRequest{
					User: &userPb.User{
						Id:       uuid.New().String(),
						Email:    "abc@gmail.com",
						UserType: enums.UserType_USER_TYPE_CUSTOMER,
						Password: "password",
					},
				},
			},
			mocks: func(md *MockDependencies) {
				md.txn.EXPECT().RunInTxn(context.Background(), gomock.Any()).Return(nil)
			},
			want: &userPb.CreateUserResponse{
				Status: rpc.StatusOk(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, md := GetServiceWithDependencies(t)
			if tt.mocks != nil {
				tt.mocks(md)
			}
			got, err := svc.CreateUser(tt.args.ctx, tt.args.req)
			if err != nil && tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			opts := []cmp.Option{
				protocmp.Transform(),
			}

			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUser(t *testing.T) {

	type args struct {
		ctx context.Context
		req *userPb.GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		mocks   func(md *MockDependencies, arg args)
		want    *userPb.GetUserResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &userPb.GetUserRequest{
					UserId: uuid.NewString(),
				},
			},
			mocks: func(md *MockDependencies, arg args) {
				md.dao.EXPECT().Get(context.Background(), gomock.Any()).Return(nil, customerrors.ErrRecordNotFound)
			},
			want: &userPb.GetUserResponse{
				Status: rpc.StatusRecordNotFound(""),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, md := GetServiceWithDependencies(t)
			if tt.mocks != nil {
				tt.mocks(md, tt.args)
			}

			got, err := svc.GetUser(tt.args.ctx, tt.args.req)
			if err != nil && tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			opts := []cmp.Option{
				protocmp.Transform(),
			}

			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
