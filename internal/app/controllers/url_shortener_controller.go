package controllers

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/request"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/response"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/internal/app/services"
	"github.com/gin-gonic/gin"
)

type UrlShortenerController struct {
	ShortenerService *services.ShortenerService
}

func NewUrlShortenerController(shortenerService *services.ShortenerService) *UrlShortenerController {
	return &UrlShortenerController{
		ShortenerService: shortenerService,
	}
}

func (usc *UrlShortenerController) ShortUrl(ctx *gin.Context) {
	setShortRequest := &request.SetShortenURLRequest{}
	err := ctx.ShouldBindJSON(setShortRequest)
	if err != nil {
		result.Error[*response.SetShortenURLResponse](err).ToJson(ctx)
		return
	}

	err = setShortRequest.Validate()
	if err != nil {
		result.BadRequest[*response.SetShortenURLResponse](err.Error()).ToJson(ctx)
		return
	}

	res := usc.ShortenerService.SetShortenedUrl(ctx, setShortRequest)
	res.ToJson(ctx)
}

func (usc *UrlShortenerController) GetShortUrl(ctx *gin.Context) {
	getShortRequest := &request.GetShortenedURLRequest{}
	err := ctx.ShouldBindUri(getShortRequest)
	if err != nil {
		result.Error[*response.GetShortenedURLResponse](err).ToJson(ctx)
		return
	}

	err = getShortRequest.Validate()
	if err != nil {
		result.BadRequest[*response.GetShortenedURLResponse](err.Error()).ToJson(ctx)
		return
	}

	res := usc.ShortenerService.GetShortenedUrl(ctx, getShortRequest)
	res.ToJson(ctx)
}
