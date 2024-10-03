package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPucblic := Router.Group("/admin/user")
	{
		userRouterPucblic.POST("/register")
	}
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
