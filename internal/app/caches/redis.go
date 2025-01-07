package caches

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r RedisCache) Exist(ctx context.Context, key string) bool {
	_, err := r.client.Get(ctx, key).Result()
	return err != nil
}

func (r RedisCache) Remove(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r RedisCache) AddList(ctx context.Context, key string, isRight bool, value interface{}) error {
	if isRight {
		return r.client.RPush(ctx, key, value).Err()
	} else {
		return r.client.LPush(ctx, key, value).Err()
	}
}

func (r RedisCache) ListPop(ctx context.Context, key string, isRight bool) (string, error) {
	if isRight {
		return r.client.RPop(ctx, key).Result()
	} else {
		return r.client.LPop(ctx, key).Result()
	}
}
