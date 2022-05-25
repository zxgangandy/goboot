package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goboot/pkg/logger"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ResponseLogger(skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, skipPath := range skipPaths {
			reg := regexp.MustCompile(skipPath)
			if reg.MatchString(path) {
				c.Next()
				return
			}
		}

		start := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		cost := time.Since(start)
		responseBody := blw.body.String()
		logger.Info(c.Request.Context(), "ResponseLog",
			zap.Int("Status", c.Writer.Status()),
			zap.String("Path", path),
			zap.String("Response", responseBody),
			zap.Duration("Cost", cost),
		)
	}
}
