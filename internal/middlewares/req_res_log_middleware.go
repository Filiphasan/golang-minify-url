package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

type ReqResLogMiddleware struct {
	r      *gin.Engine
	logger *zap.Logger
}

func NewReqResLogMiddleware(r *gin.Engine, logger *zap.Logger) *ReqResLogMiddleware {
	return &ReqResLogMiddleware{r: r, logger: logger}
}

func (m *ReqResLogMiddleware) Use() {
	m.r.Use(func(c *gin.Context) {
		start := time.Now()

		correlationId := c.GetHeader(CorrelationID)
		reqBody, _ := io.ReadAll(c.Request.Body)
		queryParams := c.Request.URL.Query()

		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		// Log Request
		m.logger.Info("Request Logging",
			zap.String("CorrelationID", correlationId),
			zap.String("Method", c.Request.Method),
			zap.String("Path", c.Request.URL.Path),
			zap.Any("QueryParams", queryParams),
			zap.Any("Headers", c.Request.Header),
			zap.String("Body", string(reqBody)),
		)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		duration := time.Since(start)

		// Log Response
		m.logger.Info("Response Logging",
			zap.String("CorrelationID", correlationId),
			zap.Int("StatusCode", c.Writer.Status()),
			zap.String("ResponseBody", blw.body.String()),
			zap.Duration("Duration", duration),
		)
	})
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Cevabı log için saklıyoruz
	return w.ResponseWriter.Write(b)
}
