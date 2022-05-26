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

func Router(profile string) *gin.Engine {

	var r = gin.New()

	if profile == utils.ProdProfile {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	} else {
		pprof.Register(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	skipPaths := []string{"/swagger/*", "/debug/*"}

	r.Use(middleware.AccessLogger(skipPaths))
	r.Use(middleware.ResponseLogger(skipPaths))
	r.Use(gin.Recovery())

	//apiV1 := r.Group("/v1")
	//apiV1.Use()

	//{
	//	// 认证相关路由
	//	apiV1.POST("/account/create_one", controller.CreateAccount)
	//	apiV1.POST("/account/create_list", controller.CreateAccounts)
	//	apiV1.POST("/account/exist_list", controller.GetExistsAccounts)
	//	apiV1.POST("/account/find_one", controller.FindAccount)
	//	apiV1.POST("/account/find_list", controller.FindAccounts)
	//	apiV1.POST("/account/has_balance", controller.HasBalance)
	//	apiV1.POST("/account/freeze", controller.Freeze)
	//	apiV1.POST("/account/unfreeze", controller.Unfreeze)
	//	apiV1.POST("/account/deposit", controller.Deposit)
	//	apiV1.POST("/account/withdraw", controller.Withdraw)
	//	apiV1.POST("/account/transfer", controller.Transfer)
	//}

	r.GET("/ping", func(c *gin.Context) {
		logger.Info(c.Request.Context(), "ping")
		rest.R.Success(c, "pong")
	})

	return r
}
