package routes

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

type HealthRoute struct {
	router           *gin.Engine
	healthController *controllers.HealthController
}

func NewHealthRoute(router *gin.Engine, healthController *controllers.HealthController) *HealthRoute {
	return &HealthRoute{
		router:           router,
		healthController: healthController,
	}
}

func (hr *HealthRoute) SetupHealthRoutes() {
	group := hr.router.Group("/api/health-check")
	group.GET("/", func(ctx *gin.Context) {
		hr.healthController.GetHealth(ctx)
	})
}
