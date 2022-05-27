package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"goboot/internal/app/middleware"
	"goboot/pkg/logger"
	"goboot/pkg/rest"
	"goboot/pkg/utils"
)

func Router(profile string, logging *logger.Config) *gin.Engine {
	var r = gin.New()

	if profile == utils.ProdProfile {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	} else {
		pprof.Register(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(middleware.AccessLogger(logging))
	r.Use(middleware.ResponseLogger(logging))
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		logger.Info(c.Request.Context(), "ping")
		rest.R.Success(c, "pong")
	})

	return r
}
