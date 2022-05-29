package dao

import (
	"context"
	"goboot/internal/app/model"
	"goboot/pkg/database"
	"gorm.io/gorm"
)

var User = NewUserDao()

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) FindById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	result := database.DB.WithContext(ctx).Where("id =? ", id).Find(&user)
	if result.Error == gorm.ErrRecordNotFound || result.RowsAffected < 1 {
		return nil, nil
	}

	return &user, result.Error
}
