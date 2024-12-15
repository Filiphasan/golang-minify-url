package setup

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp(Router *gin.Engine) *App {
	return &App{
		Router: Router,
	}
}

func (app *App) SetupApp() {
	routes.NewHealthRoute(app.Router).SetupHealthRoutes()
}
