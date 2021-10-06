package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var PSet = wire.NewSet(Controllers,NewUserController)

type Controller interface {
	Gets() map[string]gin.HandlerFunc
	Posts() map[string]gin.HandlerFunc
	Puts() map[string]gin.HandlerFunc
	Deletes() map[string]gin.HandlerFunc
}

type defaultController struct {

}

func (this *defaultController) Gets() map[string]gin.HandlerFunc {
	return nil
}

func (this *defaultController) Posts() map[string]gin.HandlerFunc {
	return nil
}

func (this *defaultController) Puts() map[string]gin.HandlerFunc {
	return nil
}

func (this *defaultController) Deletes() map[string]gin.HandlerFunc {
	return nil
}

func Controllers(controller *UserController) []Controller {
	return []Controller{controller}
}
