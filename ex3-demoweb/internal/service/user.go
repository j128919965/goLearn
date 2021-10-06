package service

import (
	"context"
	"quezr.top/demoweb/internal/data"
	"quezr.top/demoweb/internal/entity"
)

type UserService struct {
	repo *data.UserRepo
}

func NewUserService(repo *data.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (this *UserService) GetOne(ctx context.Context,id int) *entity.User{
	return this.repo.GetOne(ctx,id)
}