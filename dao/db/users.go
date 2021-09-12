package db

import (
	"database/sql"
	"example.com/hello/entity"
)

// SelectById 模拟数据库框架返回error
func SelectById(id int) (user *entity.User, err error) {
	// 模拟不存在的记录
	if id <= 1 {
		err = sql.ErrNoRows
		return
	}

	return &entity.User{
		Id:   1,
		Name: "s",
		Age:  1,
	}, nil
}
