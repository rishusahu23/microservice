package mongo

import (
	"github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/gen/api/user/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	UserCollectionName = "users"
)

type User struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string              `bson:"user_id"`
	Email     string              `bson:"email"`
	Password  string              `bson:"password"`
	UserType  enums.UserType      `bson:"user_type"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

// PrepareForInsert ensures timestamps are set when a new record is created.
func (u *User) PrepareForInsert() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

// PrepareForUpdate ensures the UpdatedAt field is refreshed during updates.
func (u *User) PrepareForUpdate() {
	u.UpdatedAt = time.Now()
}

func (u *User) ConvertToProto() *user.User {
	return &user.User{
		Id:        u.UserId,
		Email:     u.Email,
		UserType:  u.UserType,
		Password:  u.Password,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func ConvertToModel(user *user.User) *User {
	return &User{
		UserId:    user.GetId(),
		Email:     user.GetEmail(),
		Password:  user.GetPassword(),
		UserType:  user.GetUserType(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}
}
