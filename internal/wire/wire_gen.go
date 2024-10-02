// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/danhbuidcn/go_backend_api/internal/controllers"
	"github.com/danhbuidcn/go_backend_api/internal/repositories"
	"github.com/danhbuidcn/go_backend_api/internal/services"
)

// Injectors from user.wire.go:

// rename InitUserRouterHandler func to avoid duplication
func InitUserRouterHandler() (*controllers.UsersController, error) {
	iUserRepository := repositories.NewUserRepoSitory()
	iUserService := services.NewUserService(iUserRepository)
	usersController := controllers.NewUserController(iUserService)
	return usersController, nil
}
