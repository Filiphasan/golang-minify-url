package routes

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/controllers"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/gin-gonic/gin"
)

type HealthRoute struct {
	Router       *gin.Engine
	Cache        caches.Cache
	MongoContext *database.MongoContext
}

func NewHealthRoute(router *gin.Engine, cache caches.Cache, mongoContext *database.MongoContext) *HealthRoute {
	return &HealthRoute{
		Router:       router,
		Cache:        cache,
		MongoContext: mongoContext,
	}
}

func (hr *HealthRoute) SetupHealthRoutes() {
	healthController := controllers.NewHealthController(hr.Cache, hr.MongoContext)

	group := hr.Router.Group("/api/health-check")
	group.GET("/", func(context *gin.Context) {
		healthController.GetHealth(context)
	})
}
