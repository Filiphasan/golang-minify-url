package caches

import (
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

func (r RedisCache) Ping() error {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) Exist(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) Remove(key string) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) AddList(key string, isRight bool, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisCache) ListPop(key string, isRight bool) (string, error) {
	//TODO implement me
	panic("implement me")
}
