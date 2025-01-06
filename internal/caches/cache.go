package caches

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) error
}
