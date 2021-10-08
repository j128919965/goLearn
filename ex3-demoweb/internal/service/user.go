package service

import (
	"context"
	"quezr.top/demoweb/internal/data"
	"quezr.top/demoweb/internal/do"
)

type UserService struct {
	repo *data.UserRepo
}

func NewUserService(repo *data.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (this *UserService) GetOne(ctx context.Context,id int) (*do.User,error){
	return this.repo.GetOne(ctx,id)
}

func (this *UserService) AddOne(ctx context.Context,user *do.User) error {
	this.repo.AddOne(ctx,user)
	return nil
}