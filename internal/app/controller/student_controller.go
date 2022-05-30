package controller

import (
	"github.com/gin-gonic/gin"
	"goboot/internal/app/model"
	"goboot/internal/app/service"
	"goboot/pkg/baseerr"
	"goboot/pkg/logger"
	"goboot/pkg/rest"
)

// GetUser 查找学生
// @Summary 通过学生id查找学生
// @Description get student by student id
// @Tags user
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/student/get_one [post]
func GetStudent(c *gin.Context) {
	var req model.GetStudentReq
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf(ctx, "get student bind params err : %v", err)
		rest.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	result, err := service.Student.FindById(ctx, req.StudentID)
	if err != nil {
		rest.R.Error(c, baseerr.ErrInternalServer.WithDetails(err.Error()))
		return
	}

	logger.Debugf(ctx, "student=%v", result)

	rest.R.Success(c, result)
}
