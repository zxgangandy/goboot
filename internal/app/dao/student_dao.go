package dao

import (
	"context"
	"goboot/internal/app/model"
	"goboot/pkg/database"
	"gorm.io/gorm"
)

var Student = NewStudentDao()

type StudentDao struct {
}

func NewStudentDao() *StudentDao {
	return &StudentDao{}
}

func (u *StudentDao) FindById(ctx context.Context, id int64) (*model.Student, error) {
	var student model.Student
	result := database.DB.WithContext(ctx).Where("id =? ", id).Find(&student)
	if result.Error == gorm.ErrRecordNotFound || result.RowsAffected < 1 {
		return nil, nil
	}

	return &student, result.Error
}
