package dao

import (
	"database/sql"
	"example.com/hello/dao/db"
	"example.com/hello/entity"
	"example.com/hello/errors"
)

// 在很多场景下，找不到对应的记录是很正常的现象
func GetOne(id int) (user *entity.User,err error) {
	// get from redis ...
	user = nil
	err = nil
	// 不同的存储可能报不一样的错误
	if err == nil {
		return
	}
	// cache miss
	user, err = db.SelectById(id)
	// 返回包含错误码的业务错误
	if err!=nil {
		// 屏蔽数据库层的error
		if err == sql.ErrNoRows {
			return nil,errors.NotFound
		}
		// 包装错误类型
		return nil,errors.Wrap(err,"users - GetOne",500)
	}
	return
}