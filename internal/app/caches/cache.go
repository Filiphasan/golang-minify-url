package caches

import (
	"context"
	"time"
)

type Cache interface {
	Ping(ctx context.Context) error
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Exist(ctx context.Context, key string) (bool, error)
	Remove(ctx context.Context, key string) error
	AddList(ctx context.Context, key string, isRight bool, value interface{}) error
	ListPop(ctx context.Context, key string, isRight bool) (string, error)
}
