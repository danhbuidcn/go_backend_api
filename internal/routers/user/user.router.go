package user

import (
	"github.com/danhbuidcn/go_backend_api/internal/controllers"
	"github.com/danhbuidcn/go_backend_api/internal/repositories"
	"github.com/danhbuidcn/go_backend_api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) IntitUserRouter(Router *gin.RouterGroup) {
	// publish router
	// this is non-dependency
	ur := repositories.NewUserRepoSitory()
	us := services.NewUserService(ur)
	userHanderNonDenpency := controllers.NewUsersController(us)

	// WIRE go
	// Dependency Injection (DI)
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHanderNonDenpency.Register)
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
