package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) IntitAdminRouter(Router *gin.RouterGroup) {
	// publish router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
