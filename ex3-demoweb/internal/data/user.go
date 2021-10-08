package data

import (
	"context"
	"github.com/jinzhu/gorm"
	"quezr.top/demoweb/internal/data/po"
	"quezr.top/demoweb/internal/do"
	"quezr.top/demoweb/internal/errors"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo .
func NewUserRepo(db *gorm.DB) *UserRepo {
	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&po.User{})
	return &UserRepo{
		db: db,
	}
}

func (this *UserRepo) GetOne(ctx context.Context,id int) (*do.User,error) {
	var u po.User
	err := this.db.First(&u, id).Error
	// 返回包含错误码的业务错误
	if err!=nil {
		// 屏蔽数据库层的error
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound
		}
		// 包装错误类型
		return nil, errors.Wrap(err,"users - GetOne failed",500)
	}
	return &do.User{Id: u.Id,Name: u.Name,Age: u.Age},nil
}

func (this *UserRepo) AddOne(ctx context.Context,user *do.User) error {
	u := &po.User{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}
	return this.db.Create(u).Error
}