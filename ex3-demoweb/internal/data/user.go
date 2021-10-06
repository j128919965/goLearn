package data

import (
	"context"
	"quezr.top/demoweb/internal/entity"
)

type UserRepo struct {
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) GetOne(ctx context.Context,id int) *entity.User {
	return &entity.User{Name: "lzr",Age: 21}
}