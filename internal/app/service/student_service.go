package service

import (
	"context"
	"goboot/internal/app/dao"
	"goboot/internal/app/model"
	"goboot/pkg/logger"
)

var Student = NewStudentService()

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (u *StudentService) FindById(ctx context.Context, id int64) (*model.Student, error) {
	student, err := dao.Student.FindById(ctx, id)
	logger.Infof(ctx, "student=%v", student)
	return student, err
}
