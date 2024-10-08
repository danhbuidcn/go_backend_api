//go:build wireinject

package wire

import (
	"github.com/danhbuidcn/go_backend_api/internal/controllers"
	"github.com/danhbuidcn/go_backend_api/internal/repositories"
	"github.com/danhbuidcn/go_backend_api/internal/services"
	"github.com/google/wire"
)

// rename InitUserRouterHandler func to avoid duplication
func InitUserRouterHandler() (*controllers.UsersController, error) {
	wire.Build(
		repositories.NewUserRepoSitory,
		repositories.NewUserAuthRepository,
		services.NewUserService,
		controllers.NewUserController,
	)

	return new(controllers.UsersController), nil
}
