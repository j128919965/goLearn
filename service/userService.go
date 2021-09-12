package service

import (
	"example.com/hello/dao"
	"example.com/hello/entity"
	"example.com/hello/errors"
)

// SearchUser 按照id搜索用户，用户不存在为正常状况
func SearchUser(userId int) (*entity.User,error) {
	user, err := dao.GetOne(userId)
	// 用户不存在为正常状况
	if err != nil && err != errors.NotFound{
		return nil,errors.Wrap(err,"userService - search user",500)
	}
	return user,nil
}
