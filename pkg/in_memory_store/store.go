package store

import (
	"context"
	"time"
)

type InMemoryStore interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
}
