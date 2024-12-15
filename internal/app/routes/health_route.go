package routes

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

type HealthRoute struct {
	Router *gin.Engine
}

func NewHealthRoute(Router *gin.Engine) *HealthRoute {
	return &HealthRoute{
		Router: Router,
	}
}

func (hr *HealthRoute) SetupHealthRoutes() {
	healthController := controllers.NewHealthController()

	group := hr.Router.Group("/api/health-check")
	group.GET("/", func(context *gin.Context) {
		healthController.GetHealth(context)
	})
}
