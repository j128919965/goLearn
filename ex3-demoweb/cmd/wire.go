//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"quezr.top/demoweb/internal/data"
	"quezr.top/demoweb/internal/server"
	"quezr.top/demoweb/internal/service"
)

func initApp() (*gin.Engine,error){
	panic(wire.Build(server.PSet,service.ProviderSet,data.ProviderSet))
}