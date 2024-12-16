package mongo

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	userPb "github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/gen/api/user/enums"
	customerrors "github.com/rishu/microservice/pkg/errors"
	"github.com/rishu/microservice/pkg/filters"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

var (
	UserDaoMongoTS *UserDaoMongo
	userId1        = uuid.New().String()
	email1         = "abc@gmail.com"
	password       = "password"

	user1 = &userPb.User{
		Id:       userId1,
		Email:    email1,
		UserType: enums.UserType_USER_TYPE_CUSTOMER,
		Password: password,
	}
)

func createFixture(fixtures ...*userPb.User) error {
	for _, fix := range fixtures {
		if err := UserDaoMongoTS.Create(context.Background(), fix); err != nil {
			return err
		}
	}
	return nil
}

func TestUserDaoMongo_Get(t *testing.T) {
	if err := createFixture(user1); err != nil {
		panic(err)
	}
	type args struct {
		ctx     context.Context
		options []filters.FilterOption
	}
	tests := []struct {
		name    string
		args    args
		want    *userPb.User
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:     context.Background(),
				options: []filters.FilterOption{WithUserId(userId1)},
			},
			want:    user1,
			wantErr: nil,
		},
		{
			name: "success with email",
			args: args{
				ctx:     context.Background(),
				options: []filters.FilterOption{WithEmail(email1)},
			},
			want:    user1,
			wantErr: nil,
		},
		{
			name: "RNF",
			args: args{
				ctx:     context.Background(),
				options: []filters.FilterOption{WithUserId(uuid.NewString())},
			},
			want:    nil,
			wantErr: customerrors.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserDaoMongoTS.Get(tt.args.ctx, tt.args.options...)
			if err != nil || tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			opts := []cmp.Option{
				protocmp.Transform(),
				protocmp.IgnoreFields(&userPb.User{}, "created_at", "updated_at"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
