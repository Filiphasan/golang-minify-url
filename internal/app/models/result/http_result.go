package result

import (
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/gin-gonic/gin"
)

type HttpResult[T any] struct {
	Data       T      `json:"data"`
	Message    string `json:"error"`
	StatusCode int    `json:"statusCode"`
	Error      error  `json:"-"` // Error is not serialized
}

func Success[T any](data T) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    constants.SUCCESS,
		StatusCode: constants.Ok,
		Error:      nil,
	}
}

func SuccessWMessage[T any](data T, message string) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    message,
		StatusCode: constants.Ok,
		Error:      nil,
	}
}

func Created[T any](data T, message string) HttpResult[T] {
	return HttpResult[T]{
		Data:       data,
		Message:    message,
		StatusCode: constants.Created,
		Error:      nil,
	}
}

func NoContent[T any](message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.NoContent,
		Message:    message,
		Data:       zero,
		Error:      nil,
	}
}

func BadRequest[T any](message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.BadRequest,
		Message:    message,
		Data:       zero,
		Error:      nil,
	}
}

func NotFound[T any](message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.NotFound,
		Message:    message,
		Data:       zero,
		Error:      nil,
	}
}

func Failure[T any](statusCode int, message string) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: statusCode,
		Message:    message,
		Data:       zero,
		Error:      nil,
	}
}

func FailureD[T any](data T, statusCode int, message string) HttpResult[T] {
	return HttpResult[T]{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      nil,
	}
}

func Error[T any](err error) HttpResult[T] {
	var zero T
	return HttpResult[T]{
		StatusCode: constants.InternalServerError,
		Message:    constants.ERROR,
		Data:       zero,
		Error:      err,
	}
}

func (r HttpResult[T]) ToJson(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}
