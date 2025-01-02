package strategy

import (
	"context"
	"fmt"
	"github.com/google/wire"
	userPb "github.com/rishu/microservice/gen/api/user"
	store "github.com/rishu/microservice/pkg/in_memory_store"
	"google.golang.org/protobuf/encoding/protojson"
)

type Cache struct {
	redisInMemoryStore store.InMemoryStore
}

func NewCache(redisInMemoryStore store.InMemoryStore) *Cache {
	return &Cache{redisInMemoryStore: redisInMemoryStore}
}

var (
	GetUserCacheWireSet = wire.NewSet(NewCache, wire.Bind(new(GetUserStrategy), new(*Cache)))
)

func getKey(userId string) string {
	return fmt.Sprintf("%v", userId)
}
func (c *Cache) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	userStr, err := c.redisInMemoryStore.Get(ctx, getKey(request.UserId))
	if err != nil {
		return nil, err
	}
	user := &userPb.User{}
	err = protojson.Unmarshal([]byte(userStr), user)
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{
		User: user,
	}, nil
}
