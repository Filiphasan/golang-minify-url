package redis

import (
	"context"
	"fmt"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/redis/go-redis/v9"
)

func UseRedis(appConfig *configs.AppConfig, ctx context.Context) *redis.Client {
	address := fmt.Sprintf("%s:%s", appConfig.Redis.Host, appConfig.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: appConfig.Redis.Password,
		DB:       appConfig.Redis.Database,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
