package setup

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/routes"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type App struct {
	AppConfig *configs.AppConfig
	Logger    *zap.Logger
	Router    *gin.Engine
	Mongo     *mongo.Client
	Redis     *redis.Client
}

func NewApp(appConfig *configs.AppConfig, logger *zap.Logger, Router *gin.Engine, mongo *mongo.Client, redis *redis.Client) *App {
	return &App{
		AppConfig: appConfig,
		Logger:    logger,
		Router:    Router,
		Mongo:     mongo,
		Redis:     redis,
	}
}

func (app *App) SetupApp() func(ctx context.Context) {
	redisCache := caches.NewRedisCache(app.Redis)
	mongoContext := database.NewMongoContext(app.Mongo, app.AppConfig)

	routes.NewHealthRoute(app.Router, redisCache, mongoContext).SetupHealthRoutes()

	// Return a function to be deferred
	return func(ctx context.Context) {
		_ = app.Logger.Sync()
		_ = app.Redis.Close()
		_ = app.Mongo.Disconnect(ctx)
	}
}
