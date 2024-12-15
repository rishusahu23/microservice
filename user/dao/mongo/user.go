package mongo

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/gen/api/user"
	customerrors "github.com/rishu/microservice/pkg/errors"
	"github.com/rishu/microservice/user/dao"
	model "github.com/rishu/microservice/user/dao/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDaoMongo struct {
	collection *mongo.Collection
	conf       *config.Config
}

func NewUserDaoMongo(client *mongo.Client, conf *config.Config) *UserDaoMongo {
	return &UserDaoMongo{
		collection: client.Database(conf.MongoConfig.MongoDBName).Collection(model.UserCollectionName),
	}
}

var _ dao.UserDao = &UserDaoMongo{}

func (u *UserDaoMongo) Get(ctx context.Context, userId string) (*user.User, error) {
	var userModel *model.User
	filter := bson.M{
		"user_id": userId,
	}
	if err := u.collection.FindOne(ctx, filter).Decode(userModel); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}
	return userModel.ConvertToProto(), nil
}

func (u *UserDaoMongo) Create(ctx context.Context, user *user.User) error {
	userModel := model.ConvertToModel(user)
	userModel.PrepareForInsert()

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *UserDaoMongo) Update(ctx context.Context, user *user.User) error {
	userModel := model.ConvertToModel(user)
	userModel.PrepareForInsert()

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}
