package controller

import (
	"github.com/gin-gonic/gin"
	"goboot/internal/app/model"
	"goboot/internal/app/service"
	"goboot/pkg/baseerr"
	"goboot/pkg/logger"
	"goboot/pkg/rest"
)

// GetUser 查找用户
// @Summary 通过用户id查找用户
// @Description get user by user id
// @Tags user
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/user/get_one [post]
func GetUser(c *gin.Context) {
	var req model.GetUserReq
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf(ctx, "get user bind params err : %v", err)
		rest.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	result, err := service.User.FindById(ctx, req.UserID)
	if err != nil {
		rest.R.Error(c, baseerr.ErrInternalServer.WithDetails(err.Error()))
		return
	}

	logger.Infof(ctx, "user=%v", result)

	rest.R.Success(c, result)
}
