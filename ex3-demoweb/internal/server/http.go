package server

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "log"
   "quezr.top/demoweb/internal/service"
)

func NewHttpSerever(userService *service.UserService) *gin.Engine{
   engine := gin.Default()
   engine.GET("/user", func(context *gin.Context) {
      _, err := context.Writer.WriteString(fmt.Sprintf("%v",userService.GetOne(context,1)))
      if err != nil {
         log.Println("err - ",err)
      }
   })
   return engine
}