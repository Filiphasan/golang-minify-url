package controllers

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (hc *HealthController) GetHealth(ctx *gin.Context) {
	result.Success("Server is healthy", "OK").ToJson(ctx)
}
