package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"quezr.top/demoweb/internal/controller/advice"
	"quezr.top/demoweb/internal/errors"
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

func (this *defaultController) HandleResult(ctx *gin.Context, result interface{}, err error) {
	if err!=nil {
		if err,ok:=err.(*errors.BizError);ok {
			ctx.JSON(200, advice.BizFailure(err))
			return
		}
		ctx.JSON(500,advice.ErrFailure(err))
		return
	}
	ctx.JSON(200,advice.Success(result))
}



func Controllers(controller *UserController) []Controller {
	return []Controller{controller}
}
