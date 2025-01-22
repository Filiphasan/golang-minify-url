package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CorrelationID = "X-Correlation-ID"

type CorrelationMiddleware struct {
	r *gin.Engine
}

func NewCorrelationMiddleware(r *gin.Engine) *CorrelationMiddleware {
	return &CorrelationMiddleware{r: r}
}

func (m *CorrelationMiddleware) Use() {
	m.r.Use(func(c *gin.Context) {
		correlationId := c.GetHeader(CorrelationID)
		if correlationId == "" {
			correlationId = uuid.New().String()
			c.Set(CorrelationID, correlationId)
		}

		c.Next()
	})
}
