package models

import "github.com/gin-gonic/gin"

type HttpResult[T any] struct {
	Data       T      `json:"data"`
	Message    string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func Success[T any](data T, message string) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    message,
		StatusCode: 200,
	}
}

func Failure[T any](statusCode int, message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: statusCode,
		Message:    message,
		Data:       zero,
	}
}

func (r HttpResult[T]) ToJson(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}
