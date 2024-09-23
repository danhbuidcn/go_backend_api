package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) IntitProductRouter(Router *gin.RouterGroup) {
	// publish router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
	// private router
}
