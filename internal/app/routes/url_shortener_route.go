package routes

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

type UrlShortenerRoute struct {
	router *gin.Engine
	usc    *controllers.UrlShortenerController
}

func NewUrlShortenerRoute(router *gin.Engine, usc *controllers.UrlShortenerController) *UrlShortenerRoute {
	return &UrlShortenerRoute{
		router: router,
		usc:    usc,
	}
}

func (usr *UrlShortenerRoute) SetupUrlShortenerRoutes() {
	group := usr.router.Group("/api/url-shorts")
	group.POST("/", func(ctx *gin.Context) {
		usr.usc.ShortUrl(ctx)
	})

	usr.router.GET("/:token", func(ctx *gin.Context) {
		usr.usc.GetShortUrl(ctx)
	})
}
