package controllers

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/request"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/response"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/internal/app/services"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UrlShortenerController struct {
	logger           *zap.Logger
	ShortenerService *services.ShortenerService
}

func NewUrlShortenerController(logger *zap.Logger, shortenerService *services.ShortenerService) *UrlShortenerController {
	return &UrlShortenerController{
		logger:           logger,
		ShortenerService: shortenerService,
	}
}

func (usc *UrlShortenerController) ShortUrl(ctx *gin.Context) {
	const methodName = "ShortUrl"
	setShortRequest := &request.SetShortenURLRequest{}
	err := ctx.ShouldBindJSON(setShortRequest)
	if err != nil {
		usc.logger.Error("Error while binding JSON", zap.String("method", methodName), zap.Error(err))
		result.Error[*response.SetShortenURLResponse](err).ToJson(ctx)
		return
	}

	err = setShortRequest.Validate()
	if err != nil {
		usc.logger.Error("Error while validating request", zap.String("method", methodName), zap.Error(err))
		result.BadRequest[*response.SetShortenURLResponse](err.Error()).ToJson(ctx)
		return
	}

	res := usc.ShortenerService.SetShortenedUrl(ctx, setShortRequest)
	if res.Error != nil {
		usc.logger.Error("Error while setting shortened URL", zap.String("method", methodName), zap.Error(res.Error))
	}
	res.ToJson(ctx)
}

func (usc *UrlShortenerController) GetShortUrl(ctx *gin.Context) {
	const methodName = "GetShortUrl"
	getShortRequest := &request.GetShortenedURLRequest{}
	err := ctx.ShouldBindUri(getShortRequest)
	if err != nil {
		usc.logger.Error("Error while binding URI", zap.String("method", methodName), zap.Error(err))
		result.Error[*response.GetShortenedURLResponse](err).ToJson(ctx)
		return
	}

	err = getShortRequest.Validate()
	if err != nil {
		usc.logger.Error("Error while validating request", zap.String("method", methodName), zap.Error(err))
		result.BadRequest[*response.GetShortenedURLResponse](err.Error()).ToJson(ctx)
		return
	}

	res := usc.ShortenerService.GetShortenedUrl(ctx, getShortRequest)
	if res.Error != nil {
		usc.logger.Error("Error while getting shortened URL", zap.String("method", methodName), zap.Error(res.Error))
	}

	ctx.Redirect(constants.Redirect, res.Data.LongUrl)
}
