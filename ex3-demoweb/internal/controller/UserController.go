package controller

import (
	"github.com/gin-gonic/gin"
	"quezr.top/demoweb/internal/controller/dto"
	"quezr.top/demoweb/internal/do"
	"quezr.top/demoweb/internal/service"
	"strconv"
)

type UserController struct {
	defaultController
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (this *UserController) Gets() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc {
		"/user": func(context *gin.Context) {
			id,err := strconv.Atoi(context.Query("id"))
			if err != nil {
				context.String(500,err.Error())
				return
			}
			u, err := this.userService.GetOne(context, id)
			this.HandleResult(context,u,err)
		},
	}
}

func (this *UserController) Posts() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/user": func(context *gin.Context) {
			var userDto dto.CreateUserRequest
			var err error
			if err = context.Bind(&userDto) ; err!=nil{
				this.HandleResult(context,nil,err)
				return
			}
			userDo := &do.User{Id: userDto.Id,Name: userDto.Name,Age: userDto.Age}
			err = this.userService.AddOne(context, userDo)
			this.HandleResult(context,nil,err)
		},
	}
}