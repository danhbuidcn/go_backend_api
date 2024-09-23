package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) IntitUserRouter(Router *gin.RouterGroup) {
	// publish router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
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
