package routers

import (
	c "github.com/danhbuidcn/go_backend_api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/user/:id", c.NewUsersController().GetUserById)
	}
	return r
}
