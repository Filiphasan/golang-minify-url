package middlewares

import (
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExceptionMiddleware struct {
	r      *gin.Engine
	logger *zap.Logger
}

func NewExceptionMiddleware(r *gin.Engine, logger *zap.Logger) *ExceptionMiddleware {
	return &ExceptionMiddleware{r: r, logger: logger}
}

func (m *ExceptionMiddleware) Use() {
	m.r.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				correlationId := c.GetHeader(CorrelationID)
				m.logger.Error("Panic occurred", zap.Any("Error", r), zap.String(CorrelationID, correlationId))
				result.Failure[*bool](constants.InternalServerError, "Internal Server Error").ToJson(c)
			}
		}()
		c.Next()
	})
}
