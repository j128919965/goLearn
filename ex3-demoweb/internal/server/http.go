package server

import (
    "github.com/gin-gonic/gin"
    "quezr.top/demoweb/internal/controller"
)

func NewHttpSerever(controllers []controller.Controller) *gin.Engine{
   engine := gin.Default()

   for _, oneController := range controllers {
      for s, handlerFunc := range oneController.Gets() {
        engine.GET(s,handlerFunc)
      }
      for s, handlerFunc := range oneController.Posts() {
        engine.POST(s,handlerFunc)
      }
      for s, handlerFunc := range oneController.Puts() {
        engine.PUT(s,handlerFunc)
      }
      for s, handlerFunc := range oneController.Deletes() {
        engine.DELETE(s,handlerFunc)
      }
   }

   return engine
}