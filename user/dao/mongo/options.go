package mongo

import (
	"github.com/rishu/microservice/pkg/filters"
	"go.mongodb.org/mongo-driver/bson"
)

func WithUserId(userId string) filters.FilterOption {
	return filters.NewFuncMongoFilterOption(func(filter bson.M) bson.M {
		if userId == "" {
			return filter
		}
		// Add the applicationId filter
		filter["user_id"] = userId
		return filter
	})
}

func WithEmail(email string) filters.FilterOption {
	return filters.NewFuncMongoFilterOption(func(filter bson.M) bson.M {
		if email == "" {
			return filter
		}
		// Add the applicationId filter
		filter["email"] = email
		return filter
	})
}
