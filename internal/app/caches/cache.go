package caches

import "time"

type Cache interface {
	Ping() error
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) error
	Exist(key string) bool
	Remove(key string) error
	AddList(key string, isRight bool, value interface{}) error
	ListPop(key string, isRight bool) (string, error)
}
