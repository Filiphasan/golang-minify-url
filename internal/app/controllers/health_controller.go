package controllers

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/response"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	Cache        caches.Cache
	MongoContext *database.MongoContext
}

func NewHealthController(cache caches.Cache, mongoContext *database.MongoContext) *HealthController {
	return &HealthController{
		Cache:        cache,
		MongoContext: mongoContext,
	}
}

func (hc *HealthController) GetHealth(ctx *gin.Context) {
	r := &response.HealthRes{
		Healthy: true,
		Redis: response.HealthResItem{
			Healthy: true,
			Message: "OK",
		},
		Mongo: response.HealthResItem{
			Healthy: true,
			Message: "OK",
		},
	}

	redisErr := hc.Cache.Ping(ctx)
	if redisErr != nil {
		r.Healthy = false
		r.Redis.Healthy = false
		r.Redis.Message = redisErr.Error()
	}

	mongoErr := hc.MongoContext.Ping(ctx)
	if mongoErr != nil {
		r.Healthy = false
		r.Mongo.Healthy = false
		r.Mongo.Message = mongoErr.Error()
	}

	if r.Healthy {
		result.Success(r).ToJson(ctx)
	} else {
		result.FailureD(&r, constants.ServiceUnavailable, constants.FAILED).ToJson(ctx)
	}
}
