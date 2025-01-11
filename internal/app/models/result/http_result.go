package result

import (
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/gin-gonic/gin"
)

type HttpResult[T any] struct {
	Data       T      `json:"data"`
	Message    string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func Success[T any](data T, message string) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    message,
		StatusCode: constants.Ok,
	}
}

func Created[T any](data T, message string) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    message,
		StatusCode: constants.Created,
	}
}

func NoContent[T any](message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.NoContent,
		Message:    message,
		Data:       zero,
	}
}

func BadRequest[T any](message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.BadRequest,
		Message:    message,
		Data:       zero,
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
