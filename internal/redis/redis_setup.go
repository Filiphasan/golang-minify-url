package redis

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/redis/go-redis/v9"
)

func UseRedis(appConfig *configs.AppConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     appConfig.Redis.Host + ":" + appConfig.Redis.Port,
		Password: appConfig.Redis.Password,
		DB:       appConfig.Redis.Database,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
