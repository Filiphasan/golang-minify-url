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
	mongoDb := database.UseMongo(appConfig, ctx)
	redisCache := redis.UseRedis(appConfig, ctx)

	router := gin.Default()
	setupDefer := setup.NewApp(appConfig, logger.Logger, router, mongoDb, redisCache).SetupApp()
	defer setupDefer(ctx)

	logger.Logger.Info("Hello, Golang Minify URL!", zap.String("Owner", "Hasan Erdal"))
	_ = router.Run(":5001")
}
