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

const (
	traceKey = "TraceID"
	traceLen = 10
)

func AccessLogger(config *logger.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		skipPaths := config.SkipPaths
		for _, skipPath := range skipPaths {
			reg := regexp.MustCompile(skipPath)
			if reg.MatchString(path) {
				c.Next()
				return
			}
		}

		var traceId string
		headerTraceId := c.Request.Header.Get(traceKey)
		if headerTraceId == "" {
			traceId = utils.RandomString(traceLen)
		} else {
			traceId = headerTraceId
		}

		ctx := logger.WithTrace(c.Request.Context(), strings.ToLower(traceId))
		c.Request = c.Request.WithContext(ctx)

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		request := string(bodyBytes)
		if config.Desensitize {
			request = utils.MaskJsonStr(&request, config.SkipFields)
		}

		header := utils.MaskHttpHeader(c.Request.Header, []string{"Authentication"})

		logger.Info(ctx, "AccessLog",
			zap.String("Method", c.Request.Method),
			zap.String("IP", c.ClientIP()),
			zap.String("Path", path),
			zap.Any("Header", header),
			zap.String("Query", c.Request.URL.RawQuery),
			zap.String("UserAgent", c.Request.UserAgent()),
			zap.String("Request", request),
		)
		c.Next()
	}
}
