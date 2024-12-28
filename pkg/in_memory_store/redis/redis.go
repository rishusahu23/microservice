package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/wire"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/rishu/microservice/config"
	customerrors "github.com/rishu/microservice/pkg/errors"
	store "github.com/rishu/microservice/pkg/in_memory_store"
	"time"
)

type RedisInMemoryStore struct {
	redisClient *redisV9.Client
}

func NewRedisInMemoryStore(redisClient *redisV9.Client) *RedisInMemoryStore {
	return &RedisInMemoryStore{
		redisClient: redisClient,
	}
}

var (
	_            store.InMemoryStore = &RedisInMemoryStore{}
	RedisWireSet                     = wire.NewSet(NewRedisInMemoryStore, wire.Bind(new(store.InMemoryStore), new(*RedisInMemoryStore)))
)

func (r *RedisInMemoryStore) Get(ctx context.Context, key string) (string, error) {
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redisV9.Nil) {
			return "", customerrors.ErrRecordNotFound
		}
		return "", err
	}
	return res, nil
}

func (r *RedisInMemoryStore) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	if key == "" {
		return customerrors.ErrInvalidArgument
	}
	if err := r.redisClient.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func GetRedisClient(conf *config.Config) *redisV9.Client {
	return redisV9.NewClient(&redisV9.Options{
		Addr: fmt.Sprintf("%v:%v", conf.RedisConfig.Host, conf.RedisConfig.Port),
	})

}
