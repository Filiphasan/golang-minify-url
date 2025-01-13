package main

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/Filiphasan/golang-minify-url/internal/logger"
	"github.com/Filiphasan/golang-minify-url/internal/redis"
	"github.com/Filiphasan/golang-minify-url/internal/setup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	configs.LoadAppConfig()
	appConfig := configs.GetAppConfig()

	logger.UseLogger(appConfig)
	mongoDb := database.UseMongo(ctx, appConfig)
	redisCache := redis.UseRedis(ctx, appConfig)

	router := gin.Default()
	setupDefer := setup.NewApp(appConfig, logger.Logger, router, mongoDb, redisCache).Run()
	defer setupDefer(ctx)

	logger.Logger.Info("Hello, Golang Minify URL!", zap.String("Owner", "Hasan Erdal"))
	err := router.Run(":5001")
	if err != nil {
		panic(err)
	}
}
