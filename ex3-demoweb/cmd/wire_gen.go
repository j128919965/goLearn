// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"quezr.top/demoweb/internal/controller"
	"quezr.top/demoweb/internal/data"
	"quezr.top/demoweb/internal/server"
	"quezr.top/demoweb/internal/service"
)

// Injectors from wire.go:

func initApp() (*gin.Engine, error) {
	dataData, err := data.NewData()
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	v := controller.Controllers(userController)
	engine := server.NewHttpSerever(v)
	return engine, nil
}
