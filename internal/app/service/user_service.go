package service

import (
	"context"
	"goboot/internal/app/dao"
	"goboot/internal/app/model"
	"goboot/pkg/logger"
)

var User = New()

type UserService struct {
}

func New() *UserService {
	return &UserService{}
}

func (u *UserService) FindById(ctx context.Context, id int64) (*model.User, error) {
	user, err := dao.User.FindById(ctx, id)
	logger.Infof(ctx, "user=%v", user)
	return user, err
}
