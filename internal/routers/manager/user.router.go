package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) IntitUserRouter(Router *gin.RouterGroup) {
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
