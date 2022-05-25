package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goboot/pkg/logger"
	"goboot/pkg/utils"
	"io/ioutil"
	"regexp"
	"strings"
)

const traceLen = 10

func AccessLogger(skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, skipPath := range skipPaths {
			reg := regexp.MustCompile(skipPath)
			if reg.MatchString(path) {
				c.Next()
				return
			}
		}

		traceId := utils.RandomString(traceLen)
		ctx := logger.WithTrace(c.Request.Context(), strings.ToLower(traceId))
		c.Request = c.Request.WithContext(ctx)

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		logger.Info(ctx, "AccessLog",
			zap.String("Method", c.Request.Method),
			zap.String("IP", c.ClientIP()),
			zap.String("Path", path),
			zap.String("Query", c.Request.URL.RawQuery),
			zap.String("UserAgent", c.Request.UserAgent()),
			zap.String("Request", string(bodyBytes)),
		)
		c.Next()
	}
}
