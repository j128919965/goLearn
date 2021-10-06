package controller

import (
	"github.com/gin-gonic/gin"
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
			context.JSON(200,this.userService.GetOne(context,id))
		},
	}
}
