package user

import (
	"github.com/danhbuidcn/go_backend_api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) IntitUserRouter(Router *gin.RouterGroup) {
	// publish router
	// this is non-dependency
	// ur := repositories.NewUserRepoSitory()
	// us := services.NewUserService(ur)
	// userHanderNonDenpency := controllers.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()
	// WIRE go
	// Dependency Injection (DI)
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
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
